#if !defined(GPB_GRPC_PROTOCOL_ONLY) || !GPB_GRPC_PROTOCOL_ONLY
#import "Commentpost.pbrpc.h"
#import "Commentpost.pbobjc.h"
#import <ProtoRPC/ProtoRPC.h>
#import <RxLibrary/GRXWriter+Immediate.h>


@implementation CommentPost

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wobjc-designated-initializers"

// Designated initializer
- (instancetype)initWithHost:(NSString *)host callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [super initWithHost:host
                 packageName:@"pupapi"
                 serviceName:@"CommentPost"
                 callOptions:callOptions];
}

- (instancetype)initWithHost:(NSString *)host {
  return [super initWithHost:host
                 packageName:@"pupapi"
                 serviceName:@"CommentPost"];
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

#pragma mark CommentPost(CommentPostMessage) returns (CommentPostResponse)

// Deprecated methods.
- (void)commentPostWithRequest:(CommentPostMessage *)request handler:(void(^)(CommentPostResponse *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToCommentPostWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
- (GRPCProtoCall *)RPCToCommentPostWithRequest:(CommentPostMessage *)request handler:(void(^)(CommentPostResponse *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"CommentPost"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[CommentPostResponse class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
- (GRPCUnaryProtoCall *)commentPostWithMessage:(CommentPostMessage *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"CommentPost"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[CommentPostResponse class]];
}

@end
#endif
