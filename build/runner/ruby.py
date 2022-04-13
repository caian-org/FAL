# modules
from build.task import TaskRunner


class TaskRunnerRuby(TaskRunner):
    def ruby_prepare(self):
        if self._has('Gemfile') and self._has('Gemfile.lock'):
            self.run('bundle config set --local path vendor')
            self.run('bundle install')
            self.run('rm -rf lib && mkdir -p lib')
            self.run('cp -r vendor/ruby/*/gems/* lib', warn=True)

        return self
