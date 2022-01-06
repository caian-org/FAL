from os.path import join
from os.path import exists

from invoke import task

BRIGHT = ''
RESET_ALL = ''
COUNTER = 0

try:
    from colorama import init
    from colorama import Style

    init()
    BRIGHT = Style.BRIGHT
    RESET_ALL = Style.RESET_ALL

except ModuleNotFoundError:
    pass


# ~~~~
class Command:
    def __init__(self, c, pwd=None):
        self.c = c
        self.pwd = pwd

    def _has(self, f):
        if self.pwd is not None:
            return exists(join(self.pwd, f))

        return True

    def run(self, cmd):
        global COUNTER
        COUNTER += 1

        _cmd = cmd
        preffix = f'~~~ {COUNTER}. '

        print(BRIGHT)
        print(preffix, end='')

        if self.pwd is not None:
            print(f'Changing directory to "{self.pwd}"')
            print(' ' * len(preffix), end='')

            cmd = f'cd {self.pwd} && {cmd}'

        print(f'Running: {_cmd}')
        print(RESET_ALL)

        self.c.run(cmd)
        return self


# ~~~~
class NodeProject(Command):
    def npx(self, cmd):
        self.run(f'npx {cmd}')
        return self

    def serverless(self, cmd):
        self.npx(f'serverless {cmd}')
        return self


# ~~~~
class GradleProject(Command):
    def gradle_build(self):
        if self._has('build.gradle') and self._has('gradlew'):
            self.run('gradle wrapper')
            self.run('./gradlew build')

        return self


# ~~~~
class MultiLangProject(NodeProject, GradleProject):
    def build(self):
        self.gradle_build()
        return self


# ~~~~
def _base_task(c, path, sc):
    (
        MultiLangProject(c, pwd=path)
            .build()
            .serverless(sc)
    )


@task
def deploy(c, path):
    _base_task(c, path, 'deploy')


@task
def remove(c, path):
    _base_task(c, path, 'remove')
