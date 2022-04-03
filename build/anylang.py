from build.command import Command
from build.proj_dotnet import DotNetProject
from build.proj_go import GoProject
from build.proj_gradle import GradleProject
from build.proj_node import NodeProject
from build.proj_ruby import RubyProject


class Project(NodeProject, GradleProject, DotNetProject, RubyProject, GoProject):
    def prepare(self):
        return self.gradle_prepare().dotnet_prepare().ruby_prepare().npm_prepare().go_prepare()
