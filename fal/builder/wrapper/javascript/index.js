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
    __FAL_stringFuncCall: ['string', ['string']],
    __FAL_listS3Buckets: ['string', []]
  }

  const lib = Library(getSharedLibPath(), exportedFunctions)

  /* sync */
  const stringFuncCallSync = (value) => lib.__FAL_stringFuncCall(value)
  const listS3BucketsSync = () => lib.__FAL_listS3Buckets()

  /* async */
  const stringFuncCallPromise = promisify(lib.__FAL_stringFuncCall.async)
  const listS3BucketsPromise = promisify(lib.__FAL_listS3Buckets.async)

  const stringFuncCall = async (value) => await stringFuncCallPromise(value)
  const listS3Buckets = async () => await listS3BucketsPromise()

  return {
    stringFuncCall,
    stringFuncCallSync,
    listS3Buckets,
    listS3BucketsSync
  }
}

module.exports = { init }
