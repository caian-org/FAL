# standard
from os import getcwd
from os import listdir
from os.path import dirname
from os.path import isdir
from os.path import join as joinpath


def exec_func_over_dir(c, path, func):
    here = getcwd()

    for pkg in listdir(path):
        pkg_dir = joinpath(here, path, pkg)
        if isdir(pkg_dir):
            func(c, pkg_dir)
