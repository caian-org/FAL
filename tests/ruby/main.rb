require_relative '../../fal/wrappers/ruby/fal'


def main
  puts MyLib.add_and_multiplies(3)
  puts MyLib.list_s3_buckets()
end


main()
