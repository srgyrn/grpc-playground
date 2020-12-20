'use strict';

const path = require('path')
const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')

const PROTO_PATH = path.join('../pb', 'messages.proto');
const SERVER_ADDR = 'localhost:50000';
const HelloService = grpc.load(PROTO_PATH).HelloService;
const client = new HelloService(SERVER_ADDR, grpc.credentials.createInsecure());

function main() {
    client.sayHello({Name: 'Sera'}, function (error, response) {
        if (error) {
            console.log(error);
            return;
        }
        console.log(response.Message);
    })
}

main();