const { resolve, join } = require('path')
const { promisify } = require('util')

const { Library } = require('ffi-napi')

function getSharedLibPath () {
  const libPath = resolve(__dirname, '..', '..', 'shared', 'libfal')

  if (process.platform === 'win32') {
    return libPath + '.dll'
  }

  if (process.platform === 'linux') {
    return libPath + '.so'
  }

  if (process.platform === 'darwin') {
    return libPath + '.dylib'
  }

  throw new Error('Unsupported system')
}

function init () {
  const exportedFunctions = {
    __addAndMultiplies: ['int', ['int']],
    __listS3Buckets: ['void', []]
  }

  const lib = Library(getSharedLibPath(), exportedFunctions)

  /* sync */
  const addAndMultipliesSync = (value) => lib.__addAndMultiplies(value)
  const listS3BucketsSync = () => lib.__listS3Buckets()

  /* async */
  const addAndMultipliesPromise = promisify(lib.__addAndMultiplies.async)
  const listS3BucketsPromise = promisify(lib.__listS3Buckets.async)

  const addAndMultiplies = async (value) => addAndMultipliesPromise(value)
  const listS3Buckets = async () => listS3BucketsPromise()

  return { addAndMultiplies, addAndMultipliesSync, listS3Buckets, listS3BucketsSync }
}

module.exports = { init }
