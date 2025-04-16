# CONTRIBUTING.md

Anyone is welcome to contribute to this project and help make it better. 

## Makefile

All logic for this repository is located in the Makefile. Follow the makefile for the go commands.

## Commands

Run code:

```

make run # You can just run make as run is the first command which makes it default

```

Build code: This will build the got and call it gitcheck. If you want to change the filenamea of the gitcheck, change it here.

```

make build

```

Deploy/Install code: This iwll move the code to /usr/local/bin/gitcheck. If you want to change the destination change it in the Makefile

```
  make depoy
```

test: This will run golang unit tests and provide a coverage

```

make test

```



## Testing Issues

If you're attempting to contribute to get testing working you must follow the below instructions due to limitations with git. 

When cloning the repo you will need to run the following command to prepare the unit tests:

```
cd test/sub_dir_with_git_commit_required
git init
```

This was throwing an error due to the test.js file being in a sub directory. As a consequence the .git directory had to be removed

To run the unit tests:

```

make test

```
