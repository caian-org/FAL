import os
import sys

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..', '..')))

from invoke import task
from misc import Project


# ~~~~
def _base_task(c, path, sc):
    Project(c, path).prepare().serverless(sc)


# ~~~~
@task
def deploy(c, path):
    _base_task(c, path, 'deploy')


@task
def package(c, path):
    _base_task(c, path, 'package')


@task
def remove(c, path):
    Project(c, path).serverless('remove')


@task
def call(c, path):
    Project(c, path).serverless('invoke --function handler')
