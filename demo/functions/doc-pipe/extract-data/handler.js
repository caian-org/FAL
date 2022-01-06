'use strict'

const scraper = require('./scraper')

module.exports.extractData = async (event) => {
  const d = await scraper.getSomething()
  console.log(d)

  return {
    statusCode: 200,
    body: JSON.stringify(
      {
        message: 'Go Serverless v1.0! Your function executed successfully!',
        input: event
      },
      null,
      2
    )
  }
}
