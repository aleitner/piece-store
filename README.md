## Build

Build the CLI example App
```BASH
vgo build -o piece-store-cli cmd/piece-store-cli/main.go
```

Build the Server example App
```BASH
vgo build -o piece-store-server server/main.go server/utils.go
```
Server is available at `127.0.0.1:8080`


## Run Tests

```BASH
vgo test pkg/*
```

### Storage
Each shard is saved in a directory structure, the shard hash is split into chunks and the shard is stored in directories with this structure:

```
↳ <2-bytes> directory - first two bytes of shard hash used as the directory name
  ↳ <2-bytes> directory - next two bytes of the shard has used as the directory name
    ↳ <16-bytes> file - the last remaining bytes used for the filename
```

- Maximum total directories is 16 ^ 4 = 65,536
- With 100,000 files in each it would be 6,553,600,000 files total

### Cons
- Can't quickly get statistics on data stored. (Data sizes, Disk Usage)
- Limited by file system
