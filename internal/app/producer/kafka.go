package producer

import (
	"github.com/ozonmp/pay-card-api/internal/app/repo"
	"github.com/ozonmp/pay-card-api/internal/app/sender"
	"github.com/ozonmp/pay-card-api/internal/model"
	"sync"
	"time"

	"github.com/gammazero/workerpool"
)

const bufferSize = 10

type Producer interface {
	Start()
	Close()
}

type producer struct {
	n       uint64
	timeout time.Duration

	sender sender.EventSender
	events <-chan model.CardEvent

	workerPool *workerpool.WorkerPool
	repo       repo.EventRepo

	wg   *sync.WaitGroup
	done chan bool
}

func NewKafkaProducer(
	n uint64,
	sender sender.EventSender,
	events <-chan model.CardEvent,
	workerPool *workerpool.WorkerPool,
	repo repo.EventRepo,
) Producer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &producer{
		n:          n,
		sender:     sender,
		events:     events,
		workerPool: workerPool,
		repo:       repo,
		wg:         wg,
		done:       done,
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			chunk := make([]model.CardEvent, 0, bufferSize)
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					chunk = append(chunk, event)
					if len(chunk) == cap(chunk) {
						sendEvents(p, chunk)
					}
				case <-time.After(p.timeout):
					if len(chunk) > 0 {
						sendEvents(p, chunk)
					}
				case <-p.done:
					return
				}
			}
		}()
	}
}

func sendEvents(p *producer, events []model.CardEvent) {
	ids := make([]uint64, 0, len(events))
	for _, ev := range events {
		ids = append(ids, ev.ID)
	}
	if err := p.sender.Send(events); err != nil {
		unlockEvents(p, ids)
	} else {
		removeEvents(p, ids)
	}

	events = events[:0]
}

func unlockEvents(p *producer, ids []uint64) {
	p.workerPool.Submit(func() {
		p.repo.Unlock(ids)
	})
}

func removeEvents(p *producer, ids []uint64) {
	p.workerPool.Submit(func() {
		p.repo.Remove(ids)
	})
}

func (p *producer) Close() {
	close(p.done)
	p.wg.Wait()
}
