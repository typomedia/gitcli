# `gitcli` - Git Command Line Interface

[![Go Report Card](https://goreportcard.com/badge/github.com/typomedia/gitcli)](https://goreportcard.com/report/github.com/typomedia/gitcli)
[![Go Reference](https://pkg.go.dev/badge/github.com/typomedia/gitcli.svg)](https://pkg.go.dev/github.com/typomedia/gitcli)

This is a simple command line interface for git. It is written in Go and uses the native [go-git](https://github.com/go-git/go-git) library. So there is no dependency on the `git` binary. The CLI is built with [Cobra](https://github.com/spf13/cobra-cli).

`gitcli` is fully compatible with git hooks and abke to run them. Like it is also able to run git commands in a specific directory `-C`.

## Intention

This project offers a streamlined command-line interface for Git. It's meant to complement, not replace, the `git` binary. Designed for scripting and tool integration, it simplifies common Git tasks. Additionally, it provides a learning opportunity to understand how `git` works under the hood.

## Installation

Download the latest release for your platform from [Releases](https://github.com/typomedia/gitti/releases/latest)

## Usage

    gitcli [command]

## Available Commands

    checkout    Switch branches or restore working tree files
    clone       Clone a repository into a new directory
    commit      Record changes to the repository
    completion  Generate the autocompletion script for the specified shell
    help        Help about any command
    pull        Pulls changes from the remote repository
    push        Push changes to the remote repository
    reset       Reset current HEAD to the specified state
    restore     Restore working tree files
    status      Show the working tree status
    version     Show the version of gitcli

## Flags

    -h, --help   help for gitcli

## Examples

    gitcli clone https://github.com/typomedia/gitcli.git
    gitcli checkout main
    gitcli checkout -C /path/to/repo main
    gitcli pull -C /path/to/repo
    gitcli push -C /path/to/repo
    gitcli reset -C /path/to/repo --hard
    gitcli commit -C /path/to/repo -m "Commit message"
    gitcli restore -C /path/to/repo FILE/PATH
    gitcli status -C /path/to/repo

---
Copyright © 2023 Typomedia Foundation. All rights reserved.