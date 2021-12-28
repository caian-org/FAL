import platform

from ctypes import CDLL
from ctypes import c_int

from os.path import abspath
from os.path import dirname
from os.path import realpath
from os.path import join


def get_shared_lib_path():
    build_dir = abspath(join(dirname(realpath(__file__)), '..', '..', '..', 'build'))
    kernel = platform.system()

    if kernel == 'Windows':
        return join(build_dir, 'bin', 'FAL.dll')

    lib_path = join(build_dir, 'lib', 'libFAL')

    if kernel == 'Linux':
        return lib_path + '.so'

    if kernel == 'Darwin':
        return lib_path + '.dylib'

    raise Exception('Unsupported system')


def init():
    lib = CDLL(get_shared_lib_path())
    lib.addAndMultiplies.argtypes = [c_int]

    def add_and_multiples(num):
        return lib.addAndMultiplies(num)

    def list_s3_buckets():
        return lib.listS3Buckets()

    return add_and_multiples, list_s3_buckets
