
const grpc = require("@grpc/grpc-js");
const txhash_proto = require("./proto");

const {
    Worker,
    isMainThread,
    parentPort,
    workerData,
    MessageChannel,
  } = require("worker_threads");



function DigestTxHash(call, callback) {
}
function TypedDataEncoderHash(call, callback) {
}

function HashDomain(call, callback) {
}

const { port1, port2 }  = new MessageChannel();
for (let i = 0; i < 4; i++) {
    const worker = new Worker("./src/work.js", {
        workerData: i,
    });
    worker.on("message", (result) => {
      console.log(`work end:`, result);
    });
  }
  ///
  //
  var server = new grpc.Server();
  server.addService(txhash_proto.Transaction.service, {
    DigestTxHash: DigestTxHash,
    TypedDataEncoderHash: TypedDataEncoderHash,
    HasDomain: HashDomain,
  });

  server.bindAsync(
    "127.0.0.1:50001",
    grpc.ServerCredentials.createInsecure(),
    () => {
      server.start();
      console.log("grpc server started");
    }
  );