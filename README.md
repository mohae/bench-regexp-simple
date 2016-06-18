#bench-simple-regexp
A benchmark to compare using Regular Expressions with a `strings.Index` to find a single substring in a string.

For the Regular Expressions benchmarks, there are two implementations: one does a `regexp.Compile()` and uses the resulting object, the other uses a pre-compiled RegExp.

The `findWIndexBuf()` func fills a pre-allocated slice.

The string to be searched is the output of running `VBoxManage showvminfo <machinename> --machinereadable`

## Results
This is on a single-core VM running Jessie w LXDE desktop and 8GB of RAM.  The CPU is an `Intel(R) Core(TM) i5-3570K CPU @ 3.40GHz`.

```
BenchmarkFindWRegExp              100000             17057 ns/op           44480 B/op         42 allocs/op
BenchmarkFindWRegExpPrecompiled   500000              2243 ns/op              64 B/op          2 allocs/op
BenchmarkFindWIndex              2000000               824 ns/op              32 B/op          1 allocs/op
BenchmarkFindWIndexBuf           2000000               731 ns/op               0 B/op          0 allocs/op

```
