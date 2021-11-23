# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ozonmp/pay_card_api/v1/pay_card_api.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from validate import validate_pb2 as validate_dot_validate__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n)ozonmp/pay_card_api/v1/pay_card_api.proto\x12\x16ozonmp.pay_card_api.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x17validate/validate.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x97\x02\n\x04\x43\x61rd\x12\"\n\x08owner_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x07ownerId\x12.\n\x0epayment_system\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x03R\rpaymentSystem\x12 \n\x06number\x18\x03 \x01(\tB\x08\xfa\x42\x05r\x03\x98\x01\x10R\x06number\x12(\n\x0bholder_name\x18\x04 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x02R\nholderName\x12 \n\x06\x43vcCvv\x18\x05 \x01(\tB\x08\xfa\x42\x05r\x03\x98\x01\x03R\x06\x43vcCvv\x12M\n\x0f\x65xpiration_date\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x08\xfa\x42\x05\xb2\x01\x02\x08\x01R\x0e\x65xpirationDate\"\xdc\x02\n\nUpdateCard\x12\x1e\n\x08owner_id\x18\x01 \x01(\x04H\x00R\x07ownerId\x88\x01\x01\x12*\n\x0epayment_system\x18\x02 \x01(\tH\x01R\rpaymentSystem\x88\x01\x01\x12\x1b\n\x06number\x18\x03 \x01(\tH\x02R\x06number\x88\x01\x01\x12$\n\x0bholder_name\x18\x04 \x01(\tH\x03R\nholderName\x88\x01\x01\x12\x1b\n\x06\x43vcCvv\x18\x05 \x01(\tH\x04R\x06\x43vcCvv\x88\x01\x01\x12H\n\x0f\x65xpiration_date\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampH\x05R\x0e\x65xpirationDate\x88\x01\x01\x42\x0b\n\t_owner_idB\x11\n\x0f_payment_systemB\t\n\x07_numberB\x0e\n\x0c_holder_nameB\t\n\x07_CvcCvvB\x12\n\x10_expiration_date\"%\n\x13RemoveCardV1Request\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\"U\n\x11ListCardV1Request\x12\x1f\n\x06offset\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x06offset\x12\x1f\n\x05limit\x18\x02 \x01(\x04\x42\t\xfa\x42\x06\x32\x04\x18\x64(\x00R\x05limit\"\'\n\x15\x44\x65scribeCardV1Request\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\"G\n\x13\x43reateCardV1Request\x12\x30\n\x04\x63\x61rd\x18\x01 \x01(\x0b\x32\x1c.ozonmp.pay_card_api.v1.CardR\x04\x63\x61rd\"]\n\x13UpdateCardV1Request\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x36\n\x04\x63\x61rd\x18\x02 \x01(\x0b\x32\".ozonmp.pay_card_api.v1.UpdateCardR\x04\x63\x61rd\"&\n\x14\x43reateCardV1Response\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\"H\n\x12ListCardV1Response\x12\x32\n\x05\x63\x61rds\x18\x01 \x03(\x0b\x32\x1c.ozonmp.pay_card_api.v1.CardR\x05\x63\x61rds2\xd6\x04\n\x11PayCardApiService\x12}\n\nCreateCard\x12+.ozonmp.pay_card_api.v1.CreateCardV1Request\x1a,.ozonmp.pay_card_api.v1.CreateCardV1Response\"\x14\x82\xd3\xe4\x93\x02\x0e\"\t/v1/cards:\x01*\x12l\n\nUpdateCard\x12+.ozonmp.pay_card_api.v1.UpdateCardV1Request\x1a\x16.google.protobuf.Empty\"\x19\x82\xd3\xe4\x93\x02\x13\x1a\x0e/v1/cards/{id}:\x01*\x12i\n\nRemoveCard\x12+.ozonmp.pay_card_api.v1.RemoveCardV1Request\x1a\x16.google.protobuf.Empty\"\x16\x82\xd3\xe4\x93\x02\x10*\x0e/v1/cards/{id}\x12s\n\x0c\x44\x65scribeCard\x12-.ozonmp.pay_card_api.v1.DescribeCardV1Request\x1a\x1c.ozonmp.pay_card_api.v1.Card\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/v1/cards/{id}\x12t\n\x08ListCard\x12).ozonmp.pay_card_api.v1.ListCardV1Request\x1a*.ozonmp.pay_card_api.v1.ListCardV1Response\"\x11\x82\xd3\xe4\x93\x02\x0b\x12\t/v1/cardsB>Z<github.com/ozonmp/pay-card-api/pkg/pay-card-api;pay_card_apib\x06proto3')



_CARD = DESCRIPTOR.message_types_by_name['Card']
_UPDATECARD = DESCRIPTOR.message_types_by_name['UpdateCard']
_REMOVECARDV1REQUEST = DESCRIPTOR.message_types_by_name['RemoveCardV1Request']
_LISTCARDV1REQUEST = DESCRIPTOR.message_types_by_name['ListCardV1Request']
_DESCRIBECARDV1REQUEST = DESCRIPTOR.message_types_by_name['DescribeCardV1Request']
_CREATECARDV1REQUEST = DESCRIPTOR.message_types_by_name['CreateCardV1Request']
_UPDATECARDV1REQUEST = DESCRIPTOR.message_types_by_name['UpdateCardV1Request']
_CREATECARDV1RESPONSE = DESCRIPTOR.message_types_by_name['CreateCardV1Response']
_LISTCARDV1RESPONSE = DESCRIPTOR.message_types_by_name['ListCardV1Response']
Card = _reflection.GeneratedProtocolMessageType('Card', (_message.Message,), {
  'DESCRIPTOR' : _CARD,
  '__module__' : 'ozonmp.pay_card_api.v1.pay_card_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.pay_card_api.v1.Card)
  })
_sym_db.RegisterMessage(Card)

UpdateCard = _reflection.GeneratedProtocolMessageType('UpdateCard', (_message.Message,), {
  'DESCRIPTOR' : _UPDATECARD,
  '__module__' : 'ozonmp.pay_card_api.v1.pay_card_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.pay_card_api.v1.UpdateCard)
  })
_sym_db.RegisterMessage(UpdateCard)

RemoveCardV1Request = _reflection.GeneratedProtocolMessageType('RemoveCardV1Request', (_message.Message,), {
  'DESCRIPTOR' : _REMOVECARDV1REQUEST,
  '__module__' : 'ozonmp.pay_card_api.v1.pay_card_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.pay_card_api.v1.RemoveCardV1Request)
  })
_sym_db.RegisterMessage(RemoveCardV1Request)

ListCardV1Request = _reflection.GeneratedProtocolMessageType('ListCardV1Request', (_message.Message,), {
  'DESCRIPTOR' : _LISTCARDV1REQUEST,
  '__module__' : 'ozonmp.pay_card_api.v1.pay_card_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.pay_card_api.v1.ListCardV1Request)
  })
_sym_db.RegisterMessage(ListCardV1Request)

DescribeCardV1Request = _reflection.GeneratedProtocolMessageType('DescribeCardV1Request', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBECARDV1REQUEST,
  '__module__' : 'ozonmp.pay_card_api.v1.pay_card_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.pay_card_api.v1.DescribeCardV1Request)
  })
_sym_db.RegisterMessage(DescribeCardV1Request)

CreateCardV1Request = _reflection.GeneratedProtocolMessageType('CreateCardV1Request', (_message.Message,), {
  'DESCRIPTOR' : _CREATECARDV1REQUEST,
  '__module__' : 'ozonmp.pay_card_api.v1.pay_card_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.pay_card_api.v1.CreateCardV1Request)
  })
_sym_db.RegisterMessage(CreateCardV1Request)

UpdateCardV1Request = _reflection.GeneratedProtocolMessageType('UpdateCardV1Request', (_message.Message,), {
  'DESCRIPTOR' : _UPDATECARDV1REQUEST,
  '__module__' : 'ozonmp.pay_card_api.v1.pay_card_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.pay_card_api.v1.UpdateCardV1Request)
  })
_sym_db.RegisterMessage(UpdateCardV1Request)

CreateCardV1Response = _reflection.GeneratedProtocolMessageType('CreateCardV1Response', (_message.Message,), {
  'DESCRIPTOR' : _CREATECARDV1RESPONSE,
  '__module__' : 'ozonmp.pay_card_api.v1.pay_card_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.pay_card_api.v1.CreateCardV1Response)
  })
_sym_db.RegisterMessage(CreateCardV1Response)

ListCardV1Response = _reflection.GeneratedProtocolMessageType('ListCardV1Response', (_message.Message,), {
  'DESCRIPTOR' : _LISTCARDV1RESPONSE,
  '__module__' : 'ozonmp.pay_card_api.v1.pay_card_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.pay_card_api.v1.ListCardV1Response)
  })
_sym_db.RegisterMessage(ListCardV1Response)

_PAYCARDAPISERVICE = DESCRIPTOR.services_by_name['PayCardApiService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z<github.com/ozonmp/pay-card-api/pkg/pay-card-api;pay_card_api'
  _CARD.fields_by_name['owner_id']._options = None
  _CARD.fields_by_name['owner_id']._serialized_options = b'\372B\0042\002 \000'
  _CARD.fields_by_name['payment_system']._options = None
  _CARD.fields_by_name['payment_system']._serialized_options = b'\372B\004r\002\020\003'
  _CARD.fields_by_name['number']._options = None
  _CARD.fields_by_name['number']._serialized_options = b'\372B\005r\003\230\001\020'
  _CARD.fields_by_name['holder_name']._options = None
  _CARD.fields_by_name['holder_name']._serialized_options = b'\372B\004r\002\020\002'
  _CARD.fields_by_name['CvcCvv']._options = None
  _CARD.fields_by_name['CvcCvv']._serialized_options = b'\372B\005r\003\230\001\003'
  _CARD.fields_by_name['expiration_date']._options = None
  _CARD.fields_by_name['expiration_date']._serialized_options = b'\372B\005\262\001\002\010\001'
  _LISTCARDV1REQUEST.fields_by_name['offset']._options = None
  _LISTCARDV1REQUEST.fields_by_name['offset']._serialized_options = b'\372B\0042\002 \000'
  _LISTCARDV1REQUEST.fields_by_name['limit']._options = None
  _LISTCARDV1REQUEST.fields_by_name['limit']._serialized_options = b'\372B\0062\004\030d(\000'
  _PAYCARDAPISERVICE.methods_by_name['CreateCard']._options = None
  _PAYCARDAPISERVICE.methods_by_name['CreateCard']._serialized_options = b'\202\323\344\223\002\016\"\t/v1/cards:\001*'
  _PAYCARDAPISERVICE.methods_by_name['UpdateCard']._options = None
  _PAYCARDAPISERVICE.methods_by_name['UpdateCard']._serialized_options = b'\202\323\344\223\002\023\032\016/v1/cards/{id}:\001*'
  _PAYCARDAPISERVICE.methods_by_name['RemoveCard']._options = None
  _PAYCARDAPISERVICE.methods_by_name['RemoveCard']._serialized_options = b'\202\323\344\223\002\020*\016/v1/cards/{id}'
  _PAYCARDAPISERVICE.methods_by_name['DescribeCard']._options = None
  _PAYCARDAPISERVICE.methods_by_name['DescribeCard']._serialized_options = b'\202\323\344\223\002\020\022\016/v1/cards/{id}'
  _PAYCARDAPISERVICE.methods_by_name['ListCard']._options = None
  _PAYCARDAPISERVICE.methods_by_name['ListCard']._serialized_options = b'\202\323\344\223\002\013\022\t/v1/cards'
  _CARD._serialized_start=187
  _CARD._serialized_end=466
  _UPDATECARD._serialized_start=469
  _UPDATECARD._serialized_end=817
  _REMOVECARDV1REQUEST._serialized_start=819
  _REMOVECARDV1REQUEST._serialized_end=856
  _LISTCARDV1REQUEST._serialized_start=858
  _LISTCARDV1REQUEST._serialized_end=943
  _DESCRIBECARDV1REQUEST._serialized_start=945
  _DESCRIBECARDV1REQUEST._serialized_end=984
  _CREATECARDV1REQUEST._serialized_start=986
  _CREATECARDV1REQUEST._serialized_end=1057
  _UPDATECARDV1REQUEST._serialized_start=1059
  _UPDATECARDV1REQUEST._serialized_end=1152
  _CREATECARDV1RESPONSE._serialized_start=1154
  _CREATECARDV1RESPONSE._serialized_end=1192
  _LISTCARDV1RESPONSE._serialized_start=1194
  _LISTCARDV1RESPONSE._serialized_end=1266
  _PAYCARDAPISERVICE._serialized_start=1269
  _PAYCARDAPISERVICE._serialized_end=1867
# @@protoc_insertion_point(module_scope)
