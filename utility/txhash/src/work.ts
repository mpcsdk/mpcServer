import { digestTxHash, Transaction } from "@unipasswallet/transactions";
const grpc = require("@grpc/grpc-js");
const txhash_proto = require("./proto");
import { _TypedDataEncoder } from "ethers/lib/utils";
import {
  TypedDataDomain,
  TypedDataField,
} from "@ethersproject/abstract-signer";
const {
  Worker,
  isMainThread,
  parentPort,
  workerData,
} = require("worker_threads");

interface TxType {
  chainId: number;
  address: string;
  number: number;
  txs: Transaction[];
}
interface DomainType {
  domain: TypedDataDomain;
  types: Record<string, TypedDataField[]>;
  value: Record<string, any>;
}
function DigestTxHash(call: any, callback: any) {
  let txobj: TxType = JSON.parse(call.request.message);
  let hash = digestTxHash(
    txobj.chainId,
    txobj.address,
    txobj.number,
    txobj.txs
  );

  callback(null, { message: hash });
}

function TypedDataEncoderHash(call: any, callback: any) {
  let domainobj: DomainType = JSON.parse(call.request.message);

  let hash = _TypedDataEncoder.hash(
    domainobj.domain,
    domainobj.types,
    domainobj.value
  );

  callback(null, { message: hash });
}
function HashDomain(call: any, callback: any) {
  let domainobj: DomainType = JSON.parse(call.request.message);

  let has = _TypedDataEncoder.hashDomain(domainobj.domain);

  callback(null, { message: has });
}

workerData.on("message", (message: any) => {
  console.log("message:", message);
});
