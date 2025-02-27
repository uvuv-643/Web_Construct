# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: sso.proto
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
    'sso.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\tsso.proto\x12\x03sso\"2\n\x0fRegisterRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"#\n\x10RegisterResponse\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\"A\n\x0cLoginRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x10\n\x08\x61pp_uuid\x18\x03 \x01(\t\"\x1c\n\rLoginResponse\x12\x0b\n\x03jwt\x18\x01 \x01(\t\"(\n\x19GetUserPermissionsRequest\x12\x0b\n\x03jwt\x18\x01 \x01(\t\"D\n\x0fUserPermissions\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12 \n\x04\x61pps\x18\x02 \x03(\x0b\x32\x12.sso.AppPermission\"K\n\rAppPermission\x12(\n\x0bpermissions\x18\x01 \x03(\x0e\x32\x13.sso.PermissionType\x12\x10\n\x08\x61pp_uuid\x18\x02 \x01(\t*\x98\x01\n\x0ePermissionType\x12\x10\n\x0cPT_UNDEFINED\x10\x00\x12\x0b\n\x07PT_READ\x10\x01\x12\x0c\n\x08PT_WRITE\x10\x02\x12\x0e\n\nPT_EXECUTE\x10\x03\x12\r\n\tPT_DELETE\x10\x04\x12\x0c\n\x08PT_SHARE\x10\x05\x12\r\n\tPT_MANAGE\x10\x06\x12\x0f\n\x0bPT_DELEGATE\x10\x07\x12\x0c\n\x08PT_AUDIT\x10\x08\x32o\n\x04\x41uth\x12\x37\n\x08Register\x12\x14.sso.RegisterRequest\x1a\x15.sso.RegisterResponse\x12.\n\x05Login\x12\x11.sso.LoginRequest\x1a\x12.sso.LoginResponse2Y\n\x0bPermissions\x12J\n\x12GetUserPermissions\x12\x1e.sso.GetUserPermissionsRequest\x1a\x14.sso.UserPermissionsB\x0bZ\t./pkg/ssob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'sso_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\t./pkg/sso'
  _globals['_PERMISSIONTYPE']._serialized_start=394
  _globals['_PERMISSIONTYPE']._serialized_end=546
  _globals['_REGISTERREQUEST']._serialized_start=18
  _globals['_REGISTERREQUEST']._serialized_end=68
  _globals['_REGISTERRESPONSE']._serialized_start=70
  _globals['_REGISTERRESPONSE']._serialized_end=105
  _globals['_LOGINREQUEST']._serialized_start=107
  _globals['_LOGINREQUEST']._serialized_end=172
  _globals['_LOGINRESPONSE']._serialized_start=174
  _globals['_LOGINRESPONSE']._serialized_end=202
  _globals['_GETUSERPERMISSIONSREQUEST']._serialized_start=204
  _globals['_GETUSERPERMISSIONSREQUEST']._serialized_end=244
  _globals['_USERPERMISSIONS']._serialized_start=246
  _globals['_USERPERMISSIONS']._serialized_end=314
  _globals['_APPPERMISSION']._serialized_start=316
  _globals['_APPPERMISSION']._serialized_end=391
  _globals['_AUTH']._serialized_start=548
  _globals['_AUTH']._serialized_end=659
  _globals['_PERMISSIONS']._serialized_start=661
  _globals['_PERMISSIONS']._serialized_end=750
# @@protoc_insertion_point(module_scope)
