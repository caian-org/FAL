const fal = require('../../fal/wrappers/javascript')
const myLib = fal.init()

async function main () {
  for (let i = 0; i < 10; i++) {
    myLib.listS3Buckets()
  }
}

main()
