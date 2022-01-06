from os import listdir
from os.path import join
from os.path import exists

from invoke import task

CYAN = ''
DIM = ''
RESET_ALL = ''

COUNTER = 0

try:
    from colorama import init
    from colorama import Style
    from colorama import Fore

    init()
    CYAN = Fore.CYAN
    DIM = Style.DIM
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

        # ...
        print('\n' + DIM + preffix + 'Running', end='')

        if self.pwd is not None:
            print(f' @ "{self.pwd}"', end='')
            cmd = f'cd {self.pwd} && {cmd}'

        print(RESET_ALL)

        # ...
        print(CYAN + ' ' * (len(preffix) + 2) + '$ ' + _cmd + RESET_ALL + '\n')
        self.c.run(cmd, **kwargs)

        return self

    def npx(self, cmd):
        self.run(f'npx {cmd}')

        return self

    def serverless(self, cmd):
        self.npx(f'serverless {cmd}')

        return self


# ~~~~
class NodeProject(Command):
    def npm_prepare(self):
        if self._has('package.json') and self._has('package-lock.json'):
            self.run('npm install')

        return self


# ~~~~
class GradleProject(Command):
    def gradle_prepare(self):
        if self._has('build.gradle') and self._has('gradlew'):
            self.run('gradle wrapper')
            self.run('./gradlew build')

        return self


# ~~~~
class DotNetProject(Command):
    def dotnet_prepare(self):
        if self._has_ext('.fsproj') or self._has_ext('.csproj'):
            fw = '--framework netcoreapp3.1'

            self.run('dotnet restore')
            self.run(f'dotnet tool install -g Amazon.Lambda.Tools {fw}', warn=True)
            self.run(f'dotnet lambda package --configuration Release {fw} --output-package package.zip')

        return self


# ~~~~
class RubyProject(Command):
    def ruby_prepare(self):
        if self._has('Gemfile') and self._has('Gemfile.lock'):
            self.run('bundle config set --local path vendor')
            self.run('bundle install')
            self.run('rm -rf lib && mkdir -p lib')
            self.run('cp -r vendor/ruby/*/gems/* lib', warn=True)

        return self


class GoProject(Command):
    def go_prepare(self):
        if self._has('go.mod') and self._has('go.sum'):
            self.run('go get')
            self.run('GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/handler handler.go')
            self.run('upx --best --lzma bin/handler')

        return self


# ~~~~
class Project(NodeProject, GradleProject, DotNetProject, RubyProject, GoProject):
    def prepare(self):
        self.gradle_prepare().dotnet_prepare().ruby_prepare().npm_prepare().go_prepare()

        return self


# ~~~~
def _base_task(c, path, sc):
    Project(c, path).prepare().serverless(sc)


@task
def deploy(c, path):
    _base_task(c, path, 'deploy')


@task
def package(c, path):
    _base_task(c, path, 'package')


@task
def remove(c, path):
    Project(c, path).serverless('remove')
