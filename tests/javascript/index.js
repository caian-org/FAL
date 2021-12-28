const fal = require('../../fal/wrappers/javascript')
const myLib = fal.init()

async function asyncTest () {
  console.log('async main; first you should see the number, then the bucket list\n')

  myLib.listS3Buckets()
  myLib.addAndMultiplies(3).then((result) => console.log(result))
}

function syncTest () {
  console.log('sync main; first you should see the bucket list, then the number\n')

  myLib.listS3BucketsSync()
  console.log(myLib.addAndMultipliesSync(3))
}

asyncTest()
