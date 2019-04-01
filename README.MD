# Validate-commit-msg(go)
> Command line utility to validate git commit messages, golang version.

Contents
- Usage

## Usage
```
Usage: validate-commit-msg [-i] [-d] [-v] [<args>]

Options:
    -i, -install    Install this tool to .git/hooks/ as commit-msg
    -d, -debug      Debug mode
    -v, -version    Print version information and quit
    args            Commit message temp file path, use by git-commit
```

### install as git hook commit-msg
- Step 1
```
$ make
```
- Step 2: copy this tool to you PATH directory

- Step 3: change directory to you git project

- Step 4: install commit-msg hook
```
$ ./validate-commit-msg -i 
```

