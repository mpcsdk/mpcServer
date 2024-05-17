import { defaultAbiCoder, solidityPack, keccak256 ,toUtf8Bytes} from "ethers/lib/utils.js";
// import { subDigest } from "@unipasswallet/utils";
// "use strict";
// const txhash_proto = require('../src/proto');
// const grpc = require('@grpc/grpc-js');

// const msg = '{"chainId":9527,"address":"0x0DF724f89fed7Ad0955A8A149B695C279533018b","number":0,"txs":[{"_isUnipassWalletTransaction":true,"callType":0,"revertOnError":true,"gasLimit":{"type":"BigNumber","hex":"0x00"},"target":"0x9e4Ac58cfBDf5CFE0685aD034Bb5C6e26363A72a","value":{"type":"BigNumber","hex":"0x01"},"data":"0xa9059cbb000000000000000000000000752ab37a4471bf059602863f6c8225816975730e0000000000000000000000000000000000000000000000008ac7230489e80000"}]}';
// function main() {
//     var client = new txhash_proto.Transaction('localhost:50051', grpc.credentials.createInsecure());
//     client.DigestTxHash({ message: msg }, function (err, response) {
//         if (err) {
//             console.error('Error: ', err);
//         }
//         else {
//             console.log(response.message);
//         }
//     });
// }
// main();
// console.log(getAddress("0x45023fce360052cfbd96d511f34bb867715eb99a"))

function subDigest(chainId, address, hash) {
  const data =     solidityPack(
    ["bytes", "uint256", "address", "bytes32"], 
    [toUtf8Bytes("\x19\x01"), chainId, address, hash])
console.log(data)
  return keccak256(data)

}

const data = defaultAbiCoder.encode(
    [
      "uint256",
      "tuple(uint8 callType,bool revertOnError,address target,uint256 gasLimit,uint256 value,bytes data)[]",
    ],
    [0, [{
        _isUnipassWalletTransaction: true,
        callType: 0,
        revertOnError: true,
        gasLimit: {
            type: "BigNumber",
            hex: "0x0"
        },
        target: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
        value:  {
            type: "BigNumber",
            hex: "0x1"
        },
        data: "0xa9059cbb000000000000000000000000aa5c1d42f766c98089a233ce1496bce18cfac5840000000000000000000000000000000000000000000000000000000000989680"
      }]],
  )

console.log(data)
const hexstr = keccak256(data)
console.log(hexstr)
const digest = subDigest(9527, "0x4878c8FA20664D36F45F6f66Ea19F131276b8F2a", hexstr)
console.log(digest)
// //
////
// const data = defaultAbiCoder.encode(
//   [
//     "uint8",
//   ],
//   [9],
// )

// console.log(data)
// const hexstr = keccak256(data)
// console.log(hexstr)
