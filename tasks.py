from invoke import task
from build import Command


SRC_DIR = 'fal'
ARTIFACT = 'falctl'


def _build(c, flags=None):
    before = 'go run -tags _beforebuild beforebuild.go'
    build = f'go build {flags or ""} -o {ARTIFACT} .'
    move = f'mv {ARTIFACT} ..'

    return Command(c, SRC_DIR).run(before).run(build).run(move)


# ~~~~
@task
def format(c):
    Command(c).run(f'gofmt -s -w {SRC_DIR}')


@task
def build(c):
    _build(c)


@task
def release(c):
    _build(c, '-ldflags "-w -s"').unset_pwd().upx(ARTIFACT)
