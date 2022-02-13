

# Todo CLI App

## Introduction

This app is a CLI Todo list app which based on Cobra Framework.

## Usage

```
Usage:
  todocli [command]

Available Commands:
  add         Add a new todo
  completion  Generate the autocompletion script for the specified shell
  done        Mark item as done
  help        Help about any command
  list        A brief description of your command

Flags:
      --datafile string   data file to store todos (default "/Users/lilei/.tridos.json")
  -h, --help              help for todocli
  -t, --toggle            Help message for toggle

Use "todocli [command] --help" for more information about a command.
```

## How to run

1. Build source code to executable program

```
go build
```
2. Run executable program
```
./todocli [command] 
```

For example:

```
./todocli add "todo 1"
```
