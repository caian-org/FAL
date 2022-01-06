const axios = require('axios')

module.exports.getSomething = async () => {
  const res = await axios.get('https://api.frankfurter.app/latest')
  return res.data
}
