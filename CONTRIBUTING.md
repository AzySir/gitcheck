# CONTRIBUTING.md

Anyone is welcome to contribute to this project and help make it better. 

## Testing Issues

When cloning the repo due to limitations with Github you will need to run the following command to prepare the unit tests:

```
cd test/sub_dir_with_git_commit_required
git init
```

This was throwing an error due to the test.js file being in a sub directory. As a consequence the .git file had to be removed