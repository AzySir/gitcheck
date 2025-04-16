# gitcheck
A CLI tool to check what directories have lingering git changes

## Problem Statement 

After many years as a consultant on many different projects, I kept coming across an annoying issue which was on departure of a project have I update all the git repositories for a single client. This is a problem because you can have many different git repositories and you have to remember which ones you have updated that you may not have pushed any changes also. 

It was quite painful running through each directory ! 

## Solution

GitCheck was made. GitCheck will run several bash commands via Golang to check for git changes and print out a list of directories that have changes yet to be commited and what the active branch is. (Hint: If the active branch is not main you may have uncommited changes or pull requests pending!)

## Usage

GitCheck is a simple CLI tool that can be run from the command line. It does not intake a directory as a parameter now - so you will need to cd into the directory you want to check. It does however check all the subdirectories of the current directory to see if any git repositories exist.

```bash

gitcheck # Then hit enter

```

