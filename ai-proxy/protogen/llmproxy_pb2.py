# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: llmproxy.proto
# Protobuf Python Version: 5.28.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    1,
    '',
    'llmproxy.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0ellmproxy.proto\x12\x08llmproxy\x1a\x1bgoogle/protobuf/empty.proto\"8\n\nLLMRequest\x12\x0b\n\x03jwt\x18\x01 \x01(\t\x12\x0f\n\x07\x63ontent\x18\x02 \x01(\t\x12\x0c\n\x04uuid\x18\x03 \x01(\t\"7\n\x08LLMReply\x12\x0b\n\x03jwt\x18\x01 \x01(\t\x12\x10\n\x08response\x18\x02 \x01(\t\x12\x0c\n\x04uuid\x18\x03 \x01(\t2\x84\x01\n\x08LLMProxy\x12=\n\x0bSendRequest\x12\x14.llmproxy.LLMRequest\x1a\x16.google.protobuf.Empty\"\x00\x12\x39\n\tSendReply\x12\x12.llmproxy.LLMReply\x1a\x16.google.protobuf.Empty\"\x00\x42\x10Z\x0e./pkg/llmproxyb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'llmproxy_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\016./pkg/llmproxy'
  _globals['_LLMREQUEST']._serialized_start=57
  _globals['_LLMREQUEST']._serialized_end=113
  _globals['_LLMREPLY']._serialized_start=115
  _globals['_LLMREPLY']._serialized_end=170
  _globals['_LLMPROXY']._serialized_start=173
  _globals['_LLMPROXY']._serialized_end=305
# @@protoc_insertion_point(module_scope)
