const { resolve } = require('path')
const { promisify } = require('util')

const { Library } = require('ffi-napi')

function init () {
  const exportedFunctions = {
    addAndMultiplies: ['int', ['int']],
    listS3Buckets: ['void', []]
  }

  const soPath = resolve(__dirname, '..', '..', '..', 'build', 'lib', 'libFAL.dylib')
  const lib = Library(soPath, exportedFunctions)

  /* sync */
  const addAndMultipliesSync = (value) => lib.addAndMultiplies(value)
  const listS3BucketsSync = () => lib.listS3Buckets()

  /* async */
  const addAndMultipliesPromise = promisify(lib.addAndMultiplies.async)
  const listS3BucketsPromise = promisify(lib.listS3Buckets.async)

  const addAndMultiplies = async (value) => addAndMultipliesPromise(value)
  const listS3Buckets = async () => listS3BucketsPromise()

  return { addAndMultiplies, addAndMultipliesSync, listS3Buckets, listS3BucketsSync }
}

module.exports = { init }
