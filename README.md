# go-perf-tests

## File IO (Write)

Measurements per 10 iterations of each test:

Test file size: 580499010 bytes (~553 Mb)

On SSD, ~1100 Mb/sec measured write speed (disk model: APPLE SSD SM0512G, filesystem: Apple_APFS)

### Sync writes:

| # | Test name                     | Test time     | ~ Mb/sec |
|---|-------------------------------|---------------|----------|
| 1 | "File"                        | 30.697027 s   | 20       |
| 2 | "BufferedFile"                | 597.71761 ms  | 1037 *   |
| 3 | "AsyncFileWithOneWorker"      | 1.0990415 s   | 563 *    |
| 4 | "AsyncFileWithManyWorkers"    | 3.1371508 s   | 197      |

### Async writes:

| # | Test name                     | Test time     | ~ Mb/sec |
|---|-------------------------------|---------------|----------|
| 1 | "FileAsync"                   | 51.1926666 s  | 12       |
| 2 | "BufferedFileAsync"           | 955.282559 ms | 648 *    |
| 3 | "AsyncFileWithOneWorkerAsync" | 3.88582015 s  | 159      |
| 4 | "AsyncFileWithManyWorkerAsync"| 3.97423097 s  | 156      |


On HDD (A), ~850 Mb/sec measured write speed

### Sync writes:

| # | Test name                     | Test time     | ~ Mb/sec |
|---|-------------------------------|---------------|----------|
| 1 | "File"                        | 8.04336877 s  | 77       |
| 2 | "BufferedFile"                | 847.852385 ms | 730 *    |
| 3 | "AsyncFileWithOneWorker"      | 1.73917575 s  | 356 *    |
| 4 | "AsyncFileWithManyWorkers"    | 3.49826814 s  | 177      |

### Async writes:

| # | Test name                     | Test time     | ~ Mb/sec |
|---|-------------------------------|---------------|----------|
| 1 | "FileAsync"                   | 12.6254856 s  | 49       |
| 2 | "BufferedFileAsync"           | 1.88664145 s  | 328 *    |
| 3 | "AsyncFileWithOneWorkerAsync" | 6.38696083 s  | 97       |
| 4 | "AsyncFileWithManyWorkerAsync"| 6.74428174 s  | 92       |

On HDD (B), ~2400 Mb/sec measured write speed

### Sync writes:

| # | Test name                     | Test time     | ~ Mb/sec |
|---|-------------------------------|---------------|----------|
| 1 | "File"                        | 3.07002209 s  | 202      |
| 2 | "BufferedFile"                | 591.654643 ms | 1046 *   |
| 3 | "AsyncFileWithOneWorker"      | 1.29223321 s  | 479 *    |
| 4 | "AsyncFileWithManyWorkers"    | 2.62231060 s  | 236      |

### Async writes:

| # | Test name                     | Test time     | ~ Mb/sec |
|---|-------------------------------|---------------|----------|
| 1 | "FileAsync"                   | 6.900129726 s | 90       |
| 2 | "BufferedFileAsync"           | 1.423837481 s | 435 *    |
| 3 | "AsyncFileWithOneWorkerAsync" | 4.687618345 s | 132      |
| 4 | "AsyncFileWithManyWorkerAsync"| 5.572252003 s | 111      |
