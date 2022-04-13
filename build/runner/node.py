# modules
from build.task import TaskRunner


class TaskRunnerNodeJS(TaskRunner):
    def npm_prepare(self):
        if self._has('package.json') and self._has('package-lock.json'):
            self.run('npm install')

        return self
