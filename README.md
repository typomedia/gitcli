# `gitcli` - Git Command Line Interface

[![Go Report Card](https://goreportcard.com/badge/github.com/typomedia/gitcli)](https://goreportcard.com/report/github.com/typomedia/gitcli)
[![Go Reference](https://pkg.go.dev/badge/github.com/typomedia/gitcli.svg)](https://pkg.go.dev/github.com/typomedia/gitcli)

## Please do not use this project in production! It is still in development and may corrupt your repository.

This is a simple command line interface for git. It is written in Go and uses the native [go-git](https://github.com/go-git/go-git) library. So there is no dependency on the `git` binary. The CLI is built with [Cobra](https://github.com/spf13/cobra-cli).

`gitcli` is fully compatible with git hooks and able to run them. Like `git` is also able to run git commands in a specific directory `-C`.

## Intention

This project offers a streamlined command-line interface for Git. It's meant to complement, not replace, the `git` binary. Designed for scripting and tool integration, it simplifies common Git tasks. Additionally, it provides a learning opportunity to understand how `git` works under the hood.

## Installation

Download the latest release for your platform from [Releases](https://github.com/typomedia/gitti/releases/latest)

## Usage

    gitcli [command]

## Commands

    add         Add file contents to the index
    checkout    Switch branches or restore working tree files [Defaults: --force=false]
    clone       Clone a repository into a new directory
    commit      Record changes to the repository
    completion  Generate the autocompletion script for the specified shell
    help        Help about any command
    fetch       Download objects and refs from another repository
    pull        Pulls changes from the remote repository
    push        Push changes to the remote repository
    remotes     List remote names
    reset       Reset current HEAD to the specified state [Defaults: --hard=false]
    restore     Restore working tree files
    status      Show the working tree status
    version     Show the version of gitcli

## Flags

    -h, --help   help for gitcli

## Examples

    gitcli clone https://github.com/typomedia/gitcli.git
    gitcli checkout main
    gitcli pull
    gitcli push
    gitcli reset --hard
    gitcli commit -m "Commit message"
    gitcli restore FILE/PATH
    gitcli status
    gitcli add FILE1 FILE2/PATH


Every command can be used with the `-C` flag to run it in a specific directory.

    gitcli -C /path/to/repo status
    gitcli -C /path/to/repo checkout main
    gitcli -C /path/to/repo pull
    gitcli -C /path/to/repo push
    gitcli -C /path/to/repo reset --hard
    gitcli -C /path/to/repo commit -m "Commit message"
    gitcli -C /path/to/repo restore FILE/PATH
    gitcli -C /path/to/repo add FILE1 FILE2/PATH

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
Copyright © 2023 Typomedia Foundation. All rights reserved.