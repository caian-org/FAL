from ctypes import CDLL
from ctypes import c_int

from os.path import abspath
from os.path import dirname
from os.path import realpath
from os.path import join


def init():
    here = dirname(realpath(__file__))
    so_path = abspath(join(here, '..', '..', '..', 'build', 'lib', 'libFAL.dylib'))

    lib = CDLL(so_path)
    lib.addAndMultiplies.argtypes = [c_int]

    def add_and_multiples(num):
        return lib.addAndMultiplies(num)

    def list_s3_buckets():
        return lib.listS3Buckets()

    return add_and_multiples, list_s3_buckets
