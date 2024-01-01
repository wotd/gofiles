# gofiles

gofiles is small golang application that can be used to create set of files and directories.

## Usage

To create set of directories and files with default settings and in default (`/tmp`) location, just execute it:

```bash
./gofiles
```

### Getting help 
```bash
❯ ./gofiles -h
Usage of ./gofiles:
  -depth int
    	Max depth of directories (default 5)
  -dir-number int
    	Max number of directories that will be created (default 3)
  -file-number int
    	Max number of files that will be created (default 10)
  -file_size string
    	Size of files. Can be single also range, like: 10k-102M (default "10M")
  -name string
    	This string will be part of directories and files names (default "ddn")
  -path string
    	Path where directories and files will be created (default "/tmp")
```

### Parameters

| setting     | default value | note                                                              |
|-------------|---------------|-------------------------------------------------------------------|
| depth       |       5       | Max number of nesting directories level                           |
| dir-number  |       3       | Max number of subdirectories to be created for each di            |
| file-number |       10      | Max number of files to be created                                 |
| file_size   |      10M      | Size of files. Can be fixed (eg. 10k/M/G) or range (eg. 100M-23G) |
| name        |      xyz      | This name will be added to directories and files names            |
| path        |      /tmp     | Path where directories and files will be created                  |

### More advanced example

```bash
❯ ./gofiles --dir-number 5 -path /mnt/storage -depth 10
```

:exclamation: **High number of dir-number and --depth can increase time of execution** :exclamation: