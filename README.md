# go_cli
Some linux command line program clones in go (made for fun)

# mv

A simple tool to move files
### Flags

| Flag| Purpose                                              |
|-----|------------------------------------------------------|
|-n   |Prevent overwrite of existing files                   |
|-i   |Prompts user whether to overwrite each existing file  |
|-u   |Overwrites existing files only if they are newer      |
|-v   |Verbosely states each file moved   (or not moved)     |

### Example uses
```
./mv a.txt b.txt
```
Renames file a.txt to b.txt

```
./mv a.txt b/
```
Moves  file a.txt into directory b

```
./mv a/* b/
```
Moves all files in directory a into directory b

```
./mv -i a/g* b/
```
Moves all files in directory a that start with 'g' into directory b. If a file is duplicated, a user is prompted as to overwrite it or not.

# cat

Prints file contents
### Flags

| Flag| Purpose                                              |
|-----|------------------------------------------------------|
|-b   |Number nonempty output lines.                         |
|-E   |Display $ at the end of each line.                    |
|-s   |Suppress repeated empty output lines.                 |
|-T   |Display TAB characters as ^I.                         |
|-v   |Displays nonprinting characters, except for tabs and the end of line character. Uses caret notation.|
