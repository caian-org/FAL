from ctypes import CDLL
from ctypes import c_int

from os.path import abspath
from os.path import dirname
from os.path import realpath
from os.path import join


def load_fal():
    here = dirname(realpath(__file__))
    so_path = abspath(join(here, '..', '..', 'build', 'lib', 'libFAL.dylib'))

    return CDLL(so_path)


lib = load_fal()
lib.addAndMultiplies.argtypes = [c_int]


def add_and_multiples(num):
    return lib.addAndMultiplies(num)

def list_s3_buckets():
    return lib.listS3Buckets()
