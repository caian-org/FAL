# standard
import sys
import json
from os.path import abspath
from os.path import dirname
from os.path import join as joinpath

sys.path.append(abspath(joinpath(dirname(__file__), '..', '..')))

# 3rd-party
from invoke import task

# modules
from build import TaskRunnerGeneral
from build import exec_func_over_dir


def hmsg(*keys):
    help_messages = {
        'data': 'Data to be sent to the serverless function (Optional)',
        'path': 'The serverless function directory path',
        'dpath': 'Path containing one or more serverless function directories'
    }

    return {key: help_messages[key] for key in list(keys)}


@task(help=hmsg('path'))
def deploy(c, path):
    """
    Deploy a serverless function to AWS
    """
    TaskRunnerGeneral(c, path).prepare().serverless('deploy')


@task(help=hmsg('path'))
def package(c, path):
    """
    Create the serverless function deployment package
    """
    TaskRunnerGeneral(c, path).prepare().serverless('package')


@task(help=hmsg('path'))
def doctor(c, path):
    """
    Diagnose deprecation issues on serverless functions projects
    """
    TaskRunnerGeneral(c, path).serverless('doctor')


@task(help=hmsg('path'))
def remove(c, path):
    """
    Remove (undeploy) an already deployed serverless function
    """
    TaskRunnerGeneral(c, path).serverless('remove')


@task(help=hmsg('data', 'path'))
def call(c, path, data=None):
    """
    Call (invoke) a given function, sending data or not
    """
    cmd = 'invoke --function handler'

    if data is not None:
        data = json.dumps(dict(_fal=str(data)))
        cmd = f"{cmd} --data '{data}'"

    TaskRunnerGeneral(c, path).serverless(cmd)


@task(help=hmsg('dpath'))
def deploydir(c, dpath):
    """
    Iterate over all directories inside "dpath" and deploys the serverless functions
    """
    exec_func_over_dir(c, dpath, deploy)


@task(help=hmsg('dpath'))
def removedir(c, dpath):
    """
    Iterate over all directories inside "dpath" and removes/undeploys the serverless functions
    """
    exec_func_over_dir(c, dpath, remove)


@task(help=hmsg('dpath'))
def calldir(c, dpath, data):
    """
    Iterate over all directories inside "dpath" and calls/invokes the serverless functions
    """
    exec_func_over_dir(c, dpath, call)
