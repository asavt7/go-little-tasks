# Merge channels test

run benchmarks
```shell
go test -bench=. -benchmem .
```
result:
```shell
BenchmarkMergeChannelsGoroutines_little-8           1729            624401 ns/op            1172 B/op         13 allocs/op
BenchmarkMergeChannelsGoroutines_middle-8             21          48301023 ns/op           15088 B/op        119 allocs/op
BenchmarkMergeChannelsGoroutines_big-8                 3         492911553 ns/op           14896 B/op        147 allocs/op
BenchmarkMergeChannelsReflect_little-8               291           4394151 ns/op          662638 B/op      13647 allocs/op
BenchmarkMergeChannelsReflect_middle-8                 1        3400745813 ns/op        699101728 B/op  10239930 allocs/op
BenchmarkMergeChannelsReflect_big-8                    1        33347313857 ns/op       7065557176 B/op 103468523 allocs/op

```


When we try to merge too many channels (100k) reflect implementation will cause panic
```shell
panic: reflect.Select: too many cases (max 65536)
```