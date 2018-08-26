# go-perf-tests

## File IO (Write)

Measurements per 10 iterations of each test:

Test file size: 580499010 bytes (~553 Mb)

On SSD, 1100 Mb/sec measured write speed

| # | Test name                     | Test time     | Write speed (Mb/sec)|
|---|-------------------------------|---------------|---------------------|
| 1 | "File"                        | 26.575370647s | ~20.1  Mb/sec       |
| 2 | "FileAsync"                   | 45.370138542s | ~12.2  Mb/sec       |
| 3 | "BufferedFile"                | 556.05638ms   | ~994.6 Mb/sec *     |
| 4 | "BufferedFileAsync"           | 982.054624ms  | ~563.1 Mb/sec *     |
| 5 | "AsyncFileWithOneWorker"      | 1.098392675s  | ~503.6 Mb/sec *     |
| 5 | "AsyncFileWithManyWorkers"    | 3.113883248s  | ~177.6 Mb/sec       |
| 6 | "AsyncFileWithOneWorkerAsync" | 3.88572085s   | ~142.3 Mb/sec       |
| 7 | "AsyncFileWithManyWorkerAsync"| 3.869982992s  | ~142.9 Mb/sec       |
