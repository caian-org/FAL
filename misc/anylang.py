from misc.command import Command
from misc.proj_dotnet import DotNetProject
from misc.proj_go import GoProject
from misc.proj_gradle import GradleProject
from misc.proj_node import NodeProject
from misc.proj_ruby import RubyProject


class Project(NodeProject, GradleProject, DotNetProject, RubyProject, GoProject):
    def prepare(self):
        return self.gradle_prepare().dotnet_prepare().ruby_prepare().npm_prepare().go_prepare()
