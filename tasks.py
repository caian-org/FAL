from invoke import task
from misc import Command


SRC_DIR = 'fal'
ARTIFACT = 'falctl'


def _build(c, flags=None):
    b = f'go build {flags or ""} -o {ARTIFACT} .'
    m = f'mv {ARTIFACT} ..'

    return Command(c, SRC_DIR).run(b).run(m)


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
