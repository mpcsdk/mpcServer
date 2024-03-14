const hashserver = require("./hashserver")

const args = require('minimist')(process.argv.slice(2))
const url = args['url'] 
console.log(url)
hashserver.createServer(url)