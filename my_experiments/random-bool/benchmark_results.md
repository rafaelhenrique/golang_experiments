Running benchmark:


```
$ go test -bench=.

goos: linux
goarch: amd64
pkg: github.com/rafaelhenrique/golang_experiments/my_experiments/random-bool
BenchmarkBool-8          	50000000	        24.5 ns/op
BenchmarkAnotherBool-8   	  200000	      8541 ns/op
PASS
ok  	github.com/rafaelhenrique/golang_experiments/my_experiments/random-bool	3.054s
```
