from build.command import Command


class GradleProject(Command):
    def gradle_prepare(self):
        if self._has('build.gradle') and self._has('gradlew'):
            self.run('gradle wrapper')
            self.run('./gradlew build')

        return self
