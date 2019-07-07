#if !defined(GPB_GRPC_PROTOCOL_ONLY) || !GPB_GRPC_PROTOCOL_ONLY
#import "Addanimal.pbrpc.h"
#import "Addanimal.pbobjc.h"
#import <ProtoRPC/ProtoRPC.h>
#import <RxLibrary/GRXWriter+Immediate.h>


@implementation AddAnimal

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wobjc-designated-initializers"

// Designated initializer
- (instancetype)initWithHost:(NSString *)host callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [super initWithHost:host
                 packageName:@"pupapi"
                 serviceName:@"AddAnimal"
                 callOptions:callOptions];
}

- (instancetype)initWithHost:(NSString *)host {
  return [super initWithHost:host
                 packageName:@"pupapi"
                 serviceName:@"AddAnimal"];
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

#pragma mark AddAnimal(AddAnimalMessage) returns (AddAnimalResponse)

// Deprecated methods.
- (void)addAnimalWithRequest:(AddAnimalMessage *)request handler:(void(^)(AddAnimalResponse *_Nullable response, NSError *_Nullable error))handler{
  [[self RPCToAddAnimalWithRequest:request handler:handler] start];
}
// Returns a not-yet-started RPC object.
- (GRPCProtoCall *)RPCToAddAnimalWithRequest:(AddAnimalMessage *)request handler:(void(^)(AddAnimalResponse *_Nullable response, NSError *_Nullable error))handler{
  return [self RPCToMethod:@"AddAnimal"
            requestsWriter:[GRXWriter writerWithValue:request]
             responseClass:[AddAnimalResponse class]
        responsesWriteable:[GRXWriteable writeableWithSingleHandler:handler]];
}
- (GRPCUnaryProtoCall *)addAnimalWithMessage:(AddAnimalMessage *)message responseHandler:(id<GRPCProtoResponseHandler>)handler callOptions:(GRPCCallOptions *_Nullable)callOptions {
  return [self RPCToMethod:@"AddAnimal"
                   message:message
           responseHandler:handler
               callOptions:callOptions
             responseClass:[AddAnimalResponse class]];
}

@end
#endif
