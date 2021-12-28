import platform

from ctypes import CDLL
from ctypes import c_int

from os.path import abspath
from os.path import dirname
from os.path import realpath
from os.path import join


def get_ext():
    kernel = platform.system()

    if kernel == 'Linux':
        return 'so'

    if kernel == 'Windows':
        return 'dll'

    if kernel == 'Darwin':
        return 'dylib'

    raise Exception('Unsupported system')


def init():
    here = dirname(realpath(__file__))

    lib_file = 'libFAL.' + get_ext()
    lib_path = abspath(join(here, '..', '..', '..', 'build', 'lib', lib_file))

    lib = CDLL(lib_path)
    lib.addAndMultiplies.argtypes = [c_int]

    def add_and_multiples(num):
        return lib.addAndMultiplies(num)

    def list_s3_buckets():
        return lib.listS3Buckets()

    return add_and_multiples, list_s3_buckets
