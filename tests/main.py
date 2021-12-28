import sys

from os.path import abspath
from os.path import dirname
from os.path import realpath
from os.path import join


here = dirname(realpath(__file__))
wrappers_dir = abspath(join(here, '..', 'FAL', 'wrappers'))
sys.path.append(wrappers_dir)


from fal import add_and_multiples
from fal import list_s3_buckets


def main():
    print(add_and_multiples(3))
    list_s3_buckets()


if __name__ == '__main__':
    main()
