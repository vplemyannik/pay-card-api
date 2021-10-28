package retranslator

import (
	"errors"
	"github.com/bxcodec/faker/v3"
	"github.com/ozonmp/pay-card-api/internal/mocks"
	"github.com/ozonmp/pay-card-api/internal/model"
	"github.com/stretchr/testify/assert"
	"math"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestStart(t *testing.T) {

	t.Run("correct run and stop", func(t *testing.T) {
		repo, _, retranslator := setup(t, 2)

		repo.EXPECT().Lock(gomock.Any()).AnyTimes()

		retranslator.Start()
		retranslator.Close()
	})

	t.Run("correctly read all events and send", func(t *testing.T) {
		chuckSize := 10
		repo, sender, retranslator := setup(t, uint64(chuckSize))

		// arrange
		eventsCount := 100

		db := generate(eventsCount)
		offset := uint64(0)

		allRead := make(chan bool, 10)

		sendCount := int32(0)

		repo.
			EXPECT().
			Lock(gomock.Any()).
			DoAndReturn(func(size uint64) ([]model.CardEvent, error) {
				if offset >= uint64(len(db)) {
					return make([]model.CardEvent, 0), nil
				}
				chunk := db[offset : offset+size : offset+size]
				atomic.AddUint64(&offset, size)
				return chunk, nil
			}).
			AnyTimes()

		repo.EXPECT().Remove(gomock.Any()).Times(eventsCount)

		sender.
			EXPECT().
			Send(gomock.Any()).
			DoAndReturn(func(event *model.CardEvent) error {
				atomic.AddInt32(&sendCount, 1)
				if sendCount >= int32(eventsCount) {
					allRead <- true
				}
				return nil
			}).
			Times(eventsCount)

		retranslator.Start()

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-allRead:
					retranslator.Close()
					return
				}
			}
		}()

		wg.Wait()
		assert.Equal(t, int32(eventsCount), sendCount)
	})

	t.Run("correctly reprocess events again if error", func(t *testing.T) {
		chuckSize := 10
		repo, sender, retranslator := setup(t, uint64(chuckSize))

		// arrange
		eventsCount := 100

		db := generate(eventsCount)
		offset := uint64(0)

		allRead := make(chan bool, 10)

		sendCount := int32(0)

		reprocessCount := int32(0)

		repo.
			EXPECT().
			Lock(gomock.Any()).
			DoAndReturn(func(size uint64) ([]model.CardEvent, error) {
				if offset >= uint64(len(db)) {
					return make([]model.CardEvent, 0), nil
				}
				maxIndex := uint64(math.Min(float64(offset+size), float64(len(db))))
				chunk := db[offset:maxIndex:maxIndex]
				atomic.AddUint64(&offset, maxIndex-offset)
				return chunk, nil
			}).
			AnyTimes()

		repo.
			EXPECT().
			Unlock(gomock.Any()).
			DoAndReturn(func(eventIDs []uint64) error {
				atomic.AddUint64(&offset, -uint64(len(eventIDs)))
				atomic.AddInt32(&reprocessCount, 1)
				return nil
			}).AnyTimes()

		repo.EXPECT().Remove(gomock.Any()).AnyTimes()

		sender.
			EXPECT().
			Send(gomock.Any()).
			DoAndReturn(func(event *model.CardEvent) error {
				atomic.AddInt32(&sendCount, 1)
				for sendCount <= int32(eventsCount) {
					return errors.New("Error has occurred when send to kafka")
				}

				if sendCount == int32(eventsCount*2) {
					allRead <- true
				}
				return nil
			}).
			AnyTimes()

		retranslator.Start()

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-allRead:
					retranslator.Close()
					return
				}
			}
		}()

		wg.Wait()
		assert.Equal(t, int32(eventsCount*2), sendCount)
		assert.Equal(t, int32(eventsCount), reprocessCount)
	})
}

func generate(count int) []model.CardEvent {
	result := make([]model.CardEvent, 0, count)
	for i := 0; i < count; i++ {
		event := model.CardEvent{}
		faker.FakeData(&event)
		result = append(result, event)
	}

	return result
}

func setup(t *testing.T, chunkSize uint64) (*mocks.MockEventRepo, *mocks.MockEventSender, Retranslator) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	cfg := Config{
		ChannelSize:    512,
		ConsumerCount:  10,
		ConsumeSize:    chunkSize,
		ConsumeTimeout: 100 * time.Millisecond,
		ProducerCount:  2,
		WorkerCount:    2,
		Repo:           repo,
		Sender:         sender,
	}

	retranslator := NewRetranslator(cfg)

	return repo, sender, retranslator
}
