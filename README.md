# git-cleanup

![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)

A command-line interface to safely and quickly remove unwanted local git branches and to safely get your local repository back to what is on the remote repository

## Features

### Quickly remove local branches

Remove local branches using \* wildcard character pattern matching

```bash
git cleanup local feature*
```

### Safely remove local branches

Not only does the CLI list and prompt you before removing branches, but it also allows you to "Protect" branches using \* wildcard patterns

```bash
git cleanup protect master
git cleanup protect integration*
git cleanup protect release*
```

This won't prevent git or any other tool from removing these local branches, but this will prevent the git-cleanup from doing so

To remove a pattern from the protected list call unprotect with the exact pattern

```bash
git cleanup unprotect integration*
```

You can also update .git/config
```ini
[cleanup]
        protected = master,develop
        defaultbranch = master
```


### Safely reset back to the remote origin

If you want to get your local repository back to look the same as on the remote origin 


Ensure you specify which branch you want to restore to
```bash
git cleanup protect master -default
```

Running fresh-start will prompt you if necessary about un-committed changes or un-tracked files, it will remove all un-protected local branches and checkout the default branch. Hard reset the default branch to its remote origin and remove un-tracked files if you agreed to it

```bash
git cleanup fresh-start
```

## Installation

### Linux

Download Linux tar.gz from the [releases](https://github.com/paulsalmon-za/git-cleanup/releases)

Untar and copy the git-cleanup to git executable's location

To find where git is installed
```bash
whereis git
```

Copy to git location
```bash
cp git-cleanup /usr/bin
```

Give execution permissions
```bash
chmod +x /usr/bin/git-cleanup
```

### Windows

Download Windows tar.gz from the [releases](https://github.com/paulsalmon-za/git-cleanup/releases)

Unzip and copy the git-cleanup to git executable's location
 
```
c:\Program Files\git\mingw64\bin
```


## Author

Paul Salmon

* Github: [@paulsalmon-za](https://github.com/paulsalmon-za/)

## License

This software is licensed under the MIT license, see [LICENSE](./LICENSE) for more information.