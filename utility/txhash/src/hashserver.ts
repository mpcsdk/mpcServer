import { digestTxHash, Transaction } from "@unipasswallet/transactions";
const grpc = require('@grpc/grpc-js')
const txhash_proto = require('./proto')

interface TxType {
    chainId: number;
    address: string;
    number: number;
    txs: Transaction[];
}

function DigestTxHash(call: any, callback: any) {

    let txobj: TxType = JSON.parse(call.request.message)
    let hash = digestTxHash(txobj.chainId, txobj.address, txobj.number, txobj.txs);

    callback(null, { message: hash })
}

export function createServer() {
var server = new grpc.Server()
server.addService(txhash_proto.Transaction.service, { DigestTxHash: DigestTxHash })
server.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
    server.start()
    console.log('grpc server started')
})
}