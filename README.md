# datatop

[![Run Tests](https://github.com/felixfwu/datatop/actions/workflows/datatop.yml/badge.svg)](https://github.com/felixfwu/datatop/actions/workflows/datatop.yml)
[![codecov](https://codecov.io/gh/felixfwu/datatop/graph/badge.svg?token=VAC4EZSHEJ)](https://codecov.io/gh/felixfwu/datatop)
![GitHub](https://img.shields.io/github/license/felixfwu/datatop)

## A tool for finding top data.

## Usage
basic
```
# ./datatop
An open source tool for finding top data. http://github.com/felixfwu/datatop

Usage:
  datatop [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  fs          Find the directory with the most files.
  help        Help about any command

Flags:
  -h, --help      help for datatop
  -v, --version   version for datatop

Use "datatop [command] --help" for more information about a command.
```

top files of directories
```
Usage:
  datatop [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  fs          Find the directory with the most files.
  help        Help about any command

Flags:
  -h, --help      help for datatop
  -v, --version   version for datatop


# ./datatop fs
493             .git
13              .
5               cmd
4               .github
2               oracle
```