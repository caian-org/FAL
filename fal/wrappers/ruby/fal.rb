require 'ffi'
require 'os'


def _get_shared_lib_path
  here = File.dirname(__FILE__)
  lib_path = File.expand_path(File.join(here, '..', '..', 'shared', 'libfal'))

  if OS.windows?
    return lib_path + '.dll'
  end

  if OS.linux?
    return lib_path + '.so'
  end

  if OS.mac?
    return lib_path + '.dylib'
  end

  raise Exception.new 'Unsupported system'
end


module LibWrapper
  extend FFI::Library
  ffi_lib _get_shared_lib_path()

  attach_function '__addAndMultiplies', [:int], :int
  attach_function '__listS3Buckets', [], :void
end


module MyLib
  def self.add_and_multiplies(value)
    LibWrapper.__addAndMultiplies(value)
  end

  def self.list_s3_buckets
    LibWrapper.__listS3Buckets()
  end
end
