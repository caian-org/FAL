import platform

from ctypes import CDLL
from ctypes import c_char_p

from os.path import abspath
from os.path import dirname
from os.path import realpath
from os.path import join


def _get_shared_lib_path():
    kernel = platform.system()
    lib_path = abspath(join(dirname(realpath(__file__)), '..', '..', 'shared', 'libfal'))

    if kernel == 'Windows':
        return lib_path + '.dll'

    if kernel == 'Linux':
        return lib_path + '.so'

    if kernel == 'Darwin':
        return lib_path + '.dylib'

    raise Exception('Unsupported system')


def init():
    lib = CDLL(_get_shared_lib_path())

    lib.__FAL_stringFuncCall.argtypes = [c_char_p]
    lib.__FAL_stringFuncCall.restype = c_char_p

    lib.__FAL_listS3Buckets.argtypes = []
    lib.__FAL_listS3Buckets.restype = c_char_p

    def string_func_call(value):
        c_string_input = c_char_p(value.encode('UTF-8'))
        c_string_output = lib.__FAL_stringFuncCall(c_string_input)
        return c_string_output.decode('UTF-8')

    def list_s3_buckets():
        return lib.__FAL_listS3Buckets().decode('UTF-8')

    return string_func_call, list_s3_buckets
