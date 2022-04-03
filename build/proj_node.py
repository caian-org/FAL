from build.command import Command


class NodeProject(Command):
    def npm_prepare(self):
        if self._has('package.json') and self._has('package-lock.json'):
            self.run('npm install')

        return self
