const { resolve, join } = require('path')
const { promisify } = require('util')

const { Library } = require('ffi-napi')

function getFile () {
  const buildDir = resolve(__dirname, '..', '..', '..', 'build')

  if (process.platform === 'win32') {
    return join(buildDir, 'bin', 'FAL.dll')
  }

  const libPath = join(buildDir, 'lib', 'libFAL')

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
    addAndMultiplies: ['int', ['int']],
    listS3Buckets: ['void', []]
  }

  const lib = Library(getFile(), exportedFunctions)

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
