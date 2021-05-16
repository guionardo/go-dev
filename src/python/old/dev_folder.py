#!/usr/bin/python3

import datetime
import json
import os
import pathlib
import re
import sys


class DevFolder(object):

    def __init__(self):
        self.home = str(pathlib.Path.home())
        self.dev_folder = os.path.join(self.home, 'dev')
        # print('DEV=', self.dev_folder)
        self._fetch_dirs()
        self.paths = '\n'.join([dir[len(self.dev_folder)+1:]
                               for dir in self.dirs])

    def _fetch_dirs(self):
        conf_file = pathlib.Path(os.path.join(self.home, '.dev_folders.json'))
        if conf_file.exists():
            ctime = datetime.datetime.fromtimestamp(conf_file.stat().st_ctime)
            if (datetime.datetime.now()-ctime) < datetime.timedelta(minutes=15):
                try:
                    with open(str(conf_file)) as f:
                        self.dirs = json.loads(f.read())
                    return
                except Exception as exc:
                    pass
                    # print("Reconstructing corrupted file", conf_file)
            os.unlink(str(conf_file))
        paths = []
        start_level = len(self.dev_folder.split(os.path.sep))
        for root, dirs, _ in os.walk(self.dev_folder):
            if len(root.split(os.path.sep))-start_level > 1:
                continue
            for dir in dirs:
                if (dir.startswith('.') or
                    dir.startswith('_') or
                        dir in ['packages', 'env']):
                    continue
                paths.append(os.path.join(root, dir))
        self.dirs = paths
        with open(str(conf_file), 'w') as f:
            f.write(json.dumps(paths, indent=2))

    def list(self):
        print(self.paths)

    def _get_groups(self, where):
        expression = '('+''.join([f'(.*{w}*?)' for w in where])+')'
        matches = re.finditer(expression, self.paths, re.MULTILINE)
        group = None
        groups = []
        for _, match in enumerate(matches, start=1):
            if len(match.groups()) > 0:
                group = match.groups(1)[0]
                groups.append(group)
        return groups

    def _go(self, where):
        path = os.path.join(self.dev_folder, where)
        found = [p for p in self.dirs if p.startswith(path)]
        if found:
            print(found[0])

    def go(self, *where):
        groups = self._get_groups(where)
        if not groups:
            print("No dev folder located for ", where)
            return
        if len(groups) > 1:
            print("More than one dev folder is found for", where, '=', groups)
            return
        self._go(groups[0])


def show_help():
    print("Argumentos:")
    print("\tgo [termos de busca]")
    print("\tlist")


def main():
    args = sys.argv[1:]
    if args:
        dev_folder = DevFolder()

        if args[0] == 'go':
            dev_folder.go(*args[1:])
            return
        elif args[0] == 'list':
            dev_folder.list()
            return

    show_help()


if __name__ == '__main__':
    main()
