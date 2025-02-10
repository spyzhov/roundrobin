# RoundRobin

Compare different solutions of `RoundRobin`.

# benchmark

```
benchmark                                         iter      time/iter   bytes alloc        allocs
---------                                         ----      ---------   -----------        ------
BenchmarkRoundRobin_Next_sync/Chan-10         33720807    35.42 ns/op        0 B/op   0 allocs/op
BenchmarkRoundRobin_Next_sync/LL-Raw-10      421472883     2.85 ns/op        0 B/op   0 allocs/op
BenchmarkRoundRobin_Next_sync/LL-Mutex-10     77744880    15.18 ns/op        0 B/op   0 allocs/op
BenchmarkRoundRobin_Next_async/Chan-10         5181740   212.50 ns/op        0 B/op   0 allocs/op
BenchmarkRoundRobin_Next_async/LL-Raw-10     150145998     8.13 ns/op        0 B/op   0 allocs/op
BenchmarkRoundRobin_Next_async/LL-Mutex-10    39077106    30.73 ns/op        0 B/op   0 allocs/op
```

## pretty

```shell
$ go install github.com/cespare/prettybench@latest
$ go test -cpu 10 -bench . -benchmem | prettybench
```