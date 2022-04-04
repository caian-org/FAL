from os import listdir
from os.path import dirname
from os.path import isdir
from os.path import join as joinpath


def exec_func_over_dir(c, path, func):
    for pkg in listdir(path):
        pkg_dir = joinpath(dirname(__file__), path, pkg)
        if isdir(pkg_dir):
            func(c, pkg_dir)
