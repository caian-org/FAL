require 'ffi'
require 'os'


def _get_shared_lib_path
  here = File.dirname(__FILE__)
  build_dir = File.expand_path(File.join(here, '..', '..', '..', 'build'))

  if OS.windows?
    return File.join(build_dir, 'bin', 'FAL.dll')
  end

  lib_path = File.join(build_dir, 'lib', 'libFAL')

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

  attach_function 'addAndMultiplies', [:int], :int
  attach_function 'listS3Buckets', [], :void
end


module MyLib
  def self.add_and_multiplies(value)
    LibWrapper.addAndMultiplies(value)
  end

  def self.list_s3_buckets
    LibWrapper.listS3Buckets()
  end
end
