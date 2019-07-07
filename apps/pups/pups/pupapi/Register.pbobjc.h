// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: register.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers.h>
#else
 #import "GPBProtocolBuffers.h"
#endif

#if GOOGLE_PROTOBUF_OBJC_VERSION < 30002
#error This file was generated by a newer version of protoc which is incompatible with your Protocol Buffer library sources.
#endif
#if 30002 < GOOGLE_PROTOBUF_OBJC_MIN_SUPPORTED_VERSION
#error This file was generated by an older version of protoc which is incompatible with your Protocol Buffer library sources.
#endif

// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

CF_EXTERN_C_BEGIN

NS_ASSUME_NONNULL_BEGIN

#pragma mark - RegisterRoot

/**
 * Exposes the extension registry for this file.
 *
 * The base class provides:
 * @code
 *   + (GPBExtensionRegistry *)extensionRegistry;
 * @endcode
 * which is a @c GPBExtensionRegistry that includes all the extensions defined by
 * this file and all files that it depends on.
 **/
@interface RegisterRoot : GPBRootObject
@end

#pragma mark - RegisterMessage

typedef GPB_ENUM(RegisterMessage_FieldNumber) {
  RegisterMessage_FieldNumber_Email = 1,
  RegisterMessage_FieldNumber_FirstName = 2,
  RegisterMessage_FieldNumber_LastName = 3,
  RegisterMessage_FieldNumber_Zipcode = 4,
  RegisterMessage_FieldNumber_Password = 5,
};

@interface RegisterMessage : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@property(nonatomic, readwrite, copy, null_resettable) NSString *firstName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *lastName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *zipcode;

@property(nonatomic, readwrite, copy, null_resettable) NSString *password;

@end

#pragma mark - RegisterResponse

typedef GPB_ENUM(RegisterResponse_FieldNumber) {
  RegisterResponse_FieldNumber_Success = 1,
  RegisterResponse_FieldNumber_TakenEmail = 2,
  RegisterResponse_FieldNumber_InvalidPassword = 3,
  RegisterResponse_FieldNumber_InvalidFirstName = 4,
  RegisterResponse_FieldNumber_InvalidLastName = 5,
  RegisterResponse_FieldNumber_InvalidZipcode = 6,
};

@interface RegisterResponse : GPBMessage

@property(nonatomic, readwrite) BOOL success;

@property(nonatomic, readwrite) BOOL takenEmail;

@property(nonatomic, readwrite) BOOL invalidPassword;

@property(nonatomic, readwrite) BOOL invalidFirstName;

@property(nonatomic, readwrite) BOOL invalidLastName;

@property(nonatomic, readwrite) BOOL invalidZipcode;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)