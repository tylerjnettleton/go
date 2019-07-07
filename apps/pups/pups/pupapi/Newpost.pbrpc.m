#if !defined(GPB_GRPC_PROTOCOL_ONLY) || !GPB_GRPC_PROTOCOL_ONLY
#import "Newpost.pbrpc.h"
#import "Newpost.pbobjc.h"
#import <ProtoRPC/ProtoRPC.h>
#import <RxLibrary/GRXWriter+Immediate.h>


@implementation NewPost

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wobjc-designated-initializers"

// Designated initializer
- (instancetype)initWithHost:(NSString *)host callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [super initWithHost:host
                 packageName:@"pupapi"
                 serviceName:@"NewPost"
                 callOptions:callOptions];
}

- (instancetype)initWithHost:(NSString *)host {
  return [super initWithHost:host
                 packageName:@"pupapi"
                 serviceName:@"NewPost"];
}

#pragma clang diagnostic pop

// Override superclass initializer to disallow different package and service names.
- (instancetype)initWithHost:(NSString *)host
                 packageName:(NSString *)packageName
                 serviceName:(NSString *)serviceName {
  return [self initWithHost:host];
}

- (instancetype)initWithHost:(NSString *)host
                 packageName:(NSString *)packageName
                 serviceName:(NSString *)serviceName
                 callOptions:(GRPCCallOptions *)callOptions {
  return [self initWithHost:host callOptions:callOptions];
}

#pragma mark - Class Methods

+ (instancetype)serviceWithHost:(NSString *)host {
  return [[self alloc] initWithHost:host];
}

+ (instancetype)serviceWithHost:(NSString *)host callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [[self alloc] initWithHost:host callOptions:callOptions];
}

#pragma mark - Method Implementations

#pragma mark NewPost(NewPostMessage) returns (NewPostResponse)

// Deprecated methods.
- (void)newPostWithRequest:(NewPostMessage *)request handler:(void(^)(NewPostResponse *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToNewPostWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
- (GRPCProtoCall *)RPCToNewPostWithRequest:(NewPostMessage *)request handler:(void(^)(NewPostResponse *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"NewPost"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[NewPostResponse class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
- (GRPCUnaryProtoCall *)newPostWithMessage:(NewPostMessage *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"NewPost"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[NewPostResponse class]];
}

@end
#endif
