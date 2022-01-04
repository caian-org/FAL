require_relative '../../fal/wrappers/ruby/fal'

def main
  puts "--- ruby test ---\n\n"
  puts MyLib.string_func_call('Hello from Ruby!')
  puts MyLib.list_s3_buckets()
end

main()
