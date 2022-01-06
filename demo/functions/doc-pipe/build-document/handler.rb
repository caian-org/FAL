$LOAD_PATH.unshift(*Dir['./lib/**/lib'])

require 'json'
require 'builder/main'

def build_document(event:, context:)
  build_pdf()

  {
    statusCode: 200,
    body: {
      message: 'Go Serverless v1.0! Your function executed successfully!',
      input: event
    }.to_json
  }
end
