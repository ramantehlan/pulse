# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mibandDevice.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='mibandDevice.proto',
  package='main',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=b'\n\x12mibandDevice.proto\x12\x04main\"\x1a\n\nDeviceUUID\x12\x0c\n\x04UUID\x18\x01 \x01(\t\"*\n\nHeartBeats\x12\r\n\x05pulse\x18\x01 \x01(\t\x12\r\n\x05\x65rror\x18\x02 \x01(\t2G\n\x0cMibandDevice\x12\x37\n\rGetHeartBeats\x12\x10.main.DeviceUUID\x1a\x10.main.HeartBeats\"\x00\x30\x01\x62\x06proto3'
)




_DEVICEUUID = _descriptor.Descriptor(
  name='DeviceUUID',
  full_name='main.DeviceUUID',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='UUID', full_name='main.DeviceUUID.UUID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=28,
  serialized_end=54,
)


_HEARTBEATS = _descriptor.Descriptor(
  name='HeartBeats',
  full_name='main.HeartBeats',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pulse', full_name='main.HeartBeats.pulse', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='main.HeartBeats.error', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=56,
  serialized_end=98,
)

DESCRIPTOR.message_types_by_name['DeviceUUID'] = _DEVICEUUID
DESCRIPTOR.message_types_by_name['HeartBeats'] = _HEARTBEATS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DeviceUUID = _reflection.GeneratedProtocolMessageType('DeviceUUID', (_message.Message,), {
  'DESCRIPTOR' : _DEVICEUUID,
  '__module__' : 'mibandDevice_pb2'
  # @@protoc_insertion_point(class_scope:main.DeviceUUID)
  })
_sym_db.RegisterMessage(DeviceUUID)

HeartBeats = _reflection.GeneratedProtocolMessageType('HeartBeats', (_message.Message,), {
  'DESCRIPTOR' : _HEARTBEATS,
  '__module__' : 'mibandDevice_pb2'
  # @@protoc_insertion_point(class_scope:main.HeartBeats)
  })
_sym_db.RegisterMessage(HeartBeats)



_MIBANDDEVICE = _descriptor.ServiceDescriptor(
  name='MibandDevice',
  full_name='main.MibandDevice',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=100,
  serialized_end=171,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetHeartBeats',
    full_name='main.MibandDevice.GetHeartBeats',
    index=0,
    containing_service=None,
    input_type=_DEVICEUUID,
    output_type=_HEARTBEATS,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_MIBANDDEVICE)

DESCRIPTOR.services_by_name['MibandDevice'] = _MIBANDDEVICE

# @@protoc_insertion_point(module_scope)
