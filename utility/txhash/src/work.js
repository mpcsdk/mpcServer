const { digestTxHash, Transaction } = require( "@unipasswallet/transactions");
const grpc = require("@grpc/grpc-js");
const txhash_proto = require("./proto");
const { _TypedDataEncoder } = require( "ethers/lib/utils");
const {
  TypedDataDomain,
  TypedDataField,
} = require( "@ethersproject/abstract-signer");
const {
  Worker,
  isMainThread,
  parentPort,
  workerData,
} = require("worker_threads");

// interface TxType {
//   chainId: number;
//   address: string;
//   number: number;
//   txs: Transaction[];
// }
// interface DomainType {
//   domain: TypedDataDomain;
//   types: Record<string, TypedDataField[]>;
//   value: Record<string, any>;
// }
function DigestTxHash(call, callback) {
  let txobj= JSON.parse(call.request.message);
  let hash = digestTxHash(
    txobj.chainId,
    txobj.address,
    txobj.number,
    txobj.txs
  );

  callback(null, { message: hash });
}

function TypedDataEncoderHash(call, callback) {
  let domainobj = JSON.parse(call.request.message);

  let hash = _TypedDataEncoder.hash(
    domainobj.domain,
    domainobj.types,
    domainobj.value
  );

  callback(null, { message: hash });
}
function HashDomain(call, callback) {
  let domainobj = JSON.parse(call.request.message);

  let has = _TypedDataEncoder.hashDomain(domainobj.domain);

  callback(null, { message: has });
}


console.log(workerData)
workerData.on("message", (message) => {
  console.log("message:", message);
});
