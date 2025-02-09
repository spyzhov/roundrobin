# RoundRobin

Compare different solutions of `RoundRobin`.

# benchmark

```
benchmark                                     iter     time/iter   bytes alloc        allocs
---------                                     ----     ---------   -----------        ------
BenchmarkRoundRobin_Next/Chan-10          31746738   36.53 ns/op        0 B/op   0 allocs/op
BenchmarkRoundRobin_Next/LL-NoSplit-10   263551262    4.44 ns/op        0 B/op   0 allocs/op
BenchmarkRoundRobin_Next/LL-Mutex-10      79740487   15.29 ns/op        0 B/op   0 allocs/op
```

## pretty

```shell
$ go install github.com/cespare/prettybench@latest
$ go test -cpu 10 -bench . -benchmem | prettybench
```