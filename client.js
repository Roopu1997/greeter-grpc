const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = './greeter/greeter.proto';

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});

const greeter = grpc.loadPackageDefinition(packageDefinition).greeter

const client = new greeter.Greeter('localhost:7777', grpc.credentials.createInsecure());

client.greet({ name: 'Ray' }, (err, resp) => {
    if (!err) {
        console.log(resp);
        return;
    }

    console.log(err);
});
