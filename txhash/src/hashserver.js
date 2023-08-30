"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const transactions_1 = require("@unipasswallet/transactions");
const grpc = require('@grpc/grpc-js');
const txhash_proto = require('./proto');
function DigestTxHash(call, callback) {
    let txobj = JSON.parse(call.request.message);
    let hash = (0, transactions_1.digestTxHash)(txobj.chainId, txobj.address, txobj.number, txobj.txs);
    callback(null, { message: hash });
}
// function main() {
var server = new grpc.Server();
server.addService(txhash_proto.Transaction.service, { DigestTxHash: DigestTxHash });
server.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
    server.start();
    console.log('grpc server started');
});
// }
// main()
