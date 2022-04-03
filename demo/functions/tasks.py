import os
import sys
import json

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..', '..')))

from invoke import task
from misc import Project


@task
def deploy(c, path):
    Project(c, path).prepare().serverless('deploy')


@task
def package(c, path):
    Project(c, path).prepare().serverless('package')


@task
def doctor(c, path):
    Project(c, path).serverless('doctor')


@task
def remove(c, path):
    Project(c, path).serverless('remove')


@task
def call(c, path, inp):
    data = json.dumps(dict(_fal=str(inp)))
    Project(c, path).serverless(f"invoke --function handler --data '{data}'")
