//
//  ViewController.m
//  pups
//
//  Created by Tyler Nettleton on 7/6/19.
//  Copyright © 2019 Tyler Nettleton. All rights reserved.
//

#import "ViewController.h"

#import <GRPCClient/GRPCCall+ChannelArg.h>
#import <GRPCClient/GRPCCall+Tests.h>
#import <ProtoRPC/ProtoRPC.h>
#import "pupapi/Register.pbrpc.h"
#import "pupapi/Register.pbobjc.h"

@interface ViewController ()

@end

static NSString * const kHostAddress = @"localhost:7777";

@interface HLWResponseHandler : NSObject<GRPCProtoResponseHandler>

@end
// A response handler object dispatching messages to main queue
@implementation HLWResponseHandler

- (dispatch_queue_t)dispatchQueue {
    return dispatch_get_main_queue();
}

- (void)didReceiveProtoMessage:(RegisterResponse *)message {
    NSLog(@"%d", message.takenEmail);
    NSLog(@"%d", message.invalidPassword);
    NSLog(@"%d", message.success);

}

@end

@implementation ViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view.
    
//    HLWGreeter *client = [[HLWGreeter alloc] initWithHost:kHostAddress];
//
//    HLWHelloRequest *request = [HLWHelloRequest message];
//    request.name = @"Objective-C";
//
//    GRPCMutableCallOptions *options = [[GRPCMutableCallOptions alloc] init];
//    // this example does not use TLS (secure channel); use insecure channel instead
//    options.transportType = GRPCTransportTypeInsecure;
//    options.userAgentPrefix = @"HelloWorld/1.0";
//
//    GRPCUnaryProtoCall *call = [client sayHelloWithMessage:request
//                                           responseHandler:[[HLWResponseHandler alloc] init]
//                                               callOptions:options];
//
//    [call start];
    
    Register *client = [[Register alloc] initWithHost:kHostAddress];
    RegisterMessage *message = [RegisterMessage message];
    message.email = @"Tyler@tylernettleton.com";
    message.firstName = @"Tyler";
    message.lastName = @"Nettleton";
    message.zipcode = @"33308";
    message.password = @"asdf";
    
    GRPCMutableCallOptions *options = [[GRPCMutableCallOptions alloc] init];
    options.transportType = GRPCTransportTypeInsecure;
    options.userAgentPrefix = @"HelloWorld/1.0";
    
    GRPCUnaryProtoCall *call = [client registerUserWithMessage:message
                                               responseHandler:[[HLWResponseHandler alloc] init] callOptions:options];
    
    [call start];
    
}


@end
