import sys

from os.path import abspath
from os.path import dirname
from os.path import realpath
from os.path import join


def import_fal():
    here = dirname(realpath(__file__))
    wrappers_dir = abspath(join(here, '..', '..', 'FAL', 'wrappers', 'python'))
    sys.path.append(wrappers_dir)

    import fal
    return fal


def main():
    fal = import_fal()
    add_and_multiples, list_s3_buckets = fal.init()

    list_s3_buckets()
    print(add_and_multiples(3))


if __name__ == '__main__':
    main()
