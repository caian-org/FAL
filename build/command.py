from os import listdir
from os.path import join
from os.path import exists

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

    def set_pwd(self, pwd):
        self.pwd = pwd
        return self

    def unset_pwd(self):
        self.pwd = None
        return self

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
        print(CYAN + (' ' * (len(preffix) + 2)) + '$ ' + _cmd + RESET_ALL + '\n')
        self.c.run(cmd, **kwargs)

        return self

    def upx(self, file):
        self.run(f'upx --best --lzma {file}')
        return self

    def pnpx(self, cmd):
        self.run(f'pnpx {cmd}')
        return self

    def serverless(self, cmd):
        self.pnpx(f'serverless {cmd}')
        return self
