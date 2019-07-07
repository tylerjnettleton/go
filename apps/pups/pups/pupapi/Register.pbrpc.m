#if !defined(GPB_GRPC_PROTOCOL_ONLY) || !GPB_GRPC_PROTOCOL_ONLY
#import "Register.pbrpc.h"
#import "Register.pbobjc.h"
#import <ProtoRPC/ProtoRPC.h>
#import <RxLibrary/GRXWriter+Immediate.h>


@implementation Register

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wobjc-designated-initializers"

// Designated initializer
- (instancetype)initWithHost:(NSString *)host callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [super initWithHost:host
                 packageName:@"pupapi"
                 serviceName:@"Register"
                 callOptions:callOptions];
}

- (instancetype)initWithHost:(NSString *)host {
  return [super initWithHost:host
                 packageName:@"pupapi"
                 serviceName:@"Register"];
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

#pragma mark RegisterUser(RegisterMessage) returns (RegisterResponse)

// Deprecated methods.
- (void)registerUserWithRequest:(RegisterMessage *)request handler:(void(^)(RegisterResponse *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToRegisterUserWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
- (GRPCProtoCall *)RPCToRegisterUserWithRequest:(RegisterMessage *)request handler:(void(^)(RegisterResponse *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"RegisterUser"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[RegisterResponse class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
- (GRPCUnaryProtoCall *)registerUserWithMessage:(RegisterMessage *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"RegisterUser"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[RegisterResponse class]];
}

@end
#endif
