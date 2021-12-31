import sys

from os.path import abspath
from os.path import dirname
from os.path import realpath
from os.path import join


def import_fal():
    here = dirname(realpath(__file__))
    wrappers_dir = abspath(join(here, '..', '..', 'fal', 'wrappers', 'python'))
    sys.path.append(wrappers_dir)

    import fal
    return fal


def main():
    print('--- python test ---\n')

    fal = import_fal()
    string_func_call, list_s3_buckets = fal.init()

    print(string_func_call('Hello from Python!'))
    print(list_s3_buckets())


if __name__ == '__main__':
    main()
