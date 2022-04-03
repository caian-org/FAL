from build.command import Command


class DotNetProject(Command):
    def dotnet_prepare(self):
        if self._has_ext('.fsproj') or self._has_ext('.csproj'):
            fw = '--framework net6.0'

            self.run('dotnet restore')
            self.run(f'dotnet tool install -g Amazon.Lambda.Tools {fw}', warn=True)
            self.run(f'dotnet lambda package --configuration Release {fw} --output-package package.zip')

        return self
