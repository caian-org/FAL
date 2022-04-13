# modules
from build.runner import TaskRunnerDotNet
from build.runner import TaskRunnerGolang
from build.runner import TaskRunnerGradle
from build.runner import TaskRunnerNodeJS
from build.runner import TaskRunnerRuby


class TaskRunnerGeneral(TaskRunnerDotNet, TaskRunnerGolang, TaskRunnerGradle, TaskRunnerNodeJS, TaskRunnerRuby):
    def prepare(self):
        return self.gradle_prepare().dotnet_prepare().ruby_prepare().npm_prepare().go_prepare()
