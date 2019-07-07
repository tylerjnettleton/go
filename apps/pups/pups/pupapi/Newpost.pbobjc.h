// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: newpost.proto

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

#pragma mark - NewpostRoot

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
@interface NewpostRoot : GPBRootObject
@end

#pragma mark - NewPostMessage

typedef GPB_ENUM(NewPostMessage_FieldNumber) {
  NewPostMessage_FieldNumber_Title = 1,
  NewPostMessage_FieldNumber_Body = 2,
  NewPostMessage_FieldNumber_ImagesArray = 3,
};

@interface NewPostMessage : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *title;

@property(nonatomic, readwrite, copy, null_resettable) NSString *body;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<NSData*> *imagesArray;
/** The number of items in @c imagesArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger imagesArray_Count;

@end

#pragma mark - NewPostResponse

typedef GPB_ENUM(NewPostResponse_FieldNumber) {
  NewPostResponse_FieldNumber_Success = 1,
};

@interface NewPostResponse : GPBMessage

@property(nonatomic, readwrite) BOOL success;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
