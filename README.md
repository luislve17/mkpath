# mkpath

## mkdir + touch made simple with go

### Installation

* Go to the releases page and download binary.
* Move the binary to any location you want to use it locally, or include it in a path considered by your `$PATH` environment variable. `~/.local/bin` might be a good place.
* Enjoy

### Usage

Main feature consists on creating nested file in non-existent folder structure:

```bash
[19:41] @mkpath {main} ⊙ ./mkpath foo/bar/baz/text.txt
mkpath-MSG: Created: foo/bar/baz/text.txt
[19:42] @mkpath {main *} ⊙ tree foo
foo
└── bar
    └── baz
        └── text.txt

2 directories, 1 file
```

Assuming this `text.txt` file is then modified to contain important sensible info

```bash
[19:44] @mkpath {main *} ⊙ cat foo/bar/baz/text.txt 
league_of_legends_password=123-I-havent-touched-grass
```

`mkpath` will not overwrite any existing file by default

```bash
[19:44] @mkpath {main *} ⊙ ./mkpath foo/bar/baz/text.txt 
mkpath-ERR: Creating: foo/bar/baz/text.txt. File already exists.
mkpath-ERR: To force overwrite, please run using -f flag.
```

Or folder
```bash
[17] [19:45] @mkpath {main *} ⊙ ./mkpath foo/bar/baz/        
mkpath-ERR: Creating: foo/bar/baz/. File already exists.
mkpath-ERR: To force overwrite, please run using -f flag.
```

In order to achieve this, you would use `-f` flag

```bash
[19:46] @mkpath {main *} ⊙ ./mkpath -f foo/bar/baz/text.txt 
mkpath-MSG: Created: foo/bar/baz/text.txt
[19:46] @mkpath {main *} ⊙ tree foo                        
foo
└── bar
    └── baz
        └── text.txt

2 directories, 1 file
[19:46] @mkpath {main *} ⊙ cat foo/bar/baz/text.txt 
[19:47] @mkpath {main *} ⊙ # Empty
```

### To-Do list

* [ ] Create installation script for curl/wget.
* [ ] Improve logging system with colors
* [ ] Support folder overwritting into empty subfolders

### Liability

`mkpath` is currently under development, and it should be emphasize that I am not liable for any malfunctions or issues that may occur from using it. Use it at your own risk.