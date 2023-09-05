const txhash_proto = require('../src/proto')
const grpc = require('@grpc/grpc-js')
const msg = '{"chainId":9527,"address":"0x0DF724f89fed7Ad0955A8A149B695C279533018b","number":0,"txs":[{"_isUnipassWalletTransaction":true,"callType":0,"revertOnError":true,"gasLimit":{"type":"BigNumber","hex":"0x03"},"target":"0x9e4Ac58cfBDf5CFE0685aD034Bb5C6e26363A72a","value":{"type":"BigNumber","hex":"0x01"},"data":"0xa9059cbb000000000000000000000000752ab37a4471bf059602863f6c8225816975730e0000000000000000000000000000000000000000000000008ac7230489e80000"}],"txHash":"0x58816202f59de5a3e249898fe413a58cc3d0ecead33f82d288d3e7a28676eff3"}'

function main() {
    var client = new txhash_proto.Transaction('localhost:50051', grpc.credentials.createInsecure())
    client.DigestTxHash({ message: msg }, function(err, response) {
        if (err) {
            console.error('Error: ', err)
        } else {
            console.log(response.message)
        }
    })
}

main()