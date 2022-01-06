from os import listdir
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

        return False

    def _has_ext(self, e):
        if self.pwd is not None:
            for f in listdir(self.pwd):
                if f.endswith(e):
                    return True

        return False

    def run(self, cmd, **kwargs):
        global COUNTER
        COUNTER += 1

        _cmd = cmd
        preffix = f'~~~ {COUNTER}. '

        print('\n' + BRIGHT + preffix, end='')

        if self.pwd is not None:
            print(f'Changing directory to "{self.pwd}"')
            print(' ' * len(preffix), end='')

            cmd = f'cd {self.pwd} && {cmd}'

        print('Running:')
        print(' ' * (len(preffix) + 2), end='')
        print(cmd + RESET_ALL + '\n')

        self.c.run(cmd, **kwargs)
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
class DotNetProject(Command):
    def dotnet_build(self):
        if self._has_ext('.fsproj'):
            self.run('dotnet restore')
            self.run('dotnet tool install -g Amazon.Lambda.Tools --framework netcoreapp3.1', warn=True)
            self.run('dotnet lambda package --configuration Release --framework netcoreapp3.1 --output-package package.zip')

        return self


# ~~~~
class MultiLangProject(NodeProject, GradleProject, DotNetProject):
    def build(self):
        self.gradle_build()
        self.dotnet_build()

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
    MultiLangProject(c, pwd=path).serverless('remove')
