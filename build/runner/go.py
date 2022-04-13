# modules
from build.task import TaskRunner


class TaskRunnerGolang(TaskRunner):
    def go_prepare(self):
        if self._has('go.mod') and self._has('go.sum'):
            artf = 'bin/handler'

            self.run('go get')
            self.run(f'GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o {artf} handler.go')
            self.upx(artf)

        return self
