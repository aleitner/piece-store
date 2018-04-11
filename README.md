## Build

```BASH

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
