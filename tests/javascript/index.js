const fal = require('../../fal/wrappers/javascript')
const myLib = fal.init()

async function asyncTest () {
  console.log('async main\n'.toUpperCase())

  const a = await myLib.stringFuncCall("Hello from JS")
  const b = await myLib.listS3Buckets()

  console.log(a)
  console.log(b)
}

function syncTest () {
  console.log('sync main\n'.toUpperCase())

  console.log(myLib.listS3BucketsSync())
  console.log(myLib.stringFuncCallSync("Hello from JS"))
}

console.log('--- javascript test ---\n')

asyncTest()
  .then(() => syncTest())
