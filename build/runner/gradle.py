# modules
from build.task import TaskRunner


class TaskRunnerGradle(TaskRunner):
    def gradle_prepare(self):
        if self._has('build.gradle') and self._has('gradlew'):
            self.run('gradle wrapper')
            self.run('./gradlew build')

        return self
