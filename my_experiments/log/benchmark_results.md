
```
$ go test -bench=.

goos: linux
goarch: amd64
pkg: github.com/rafaelhenrique/golang_experiments/my_experiments/log
BenchmarkLog-8                               	 1000000	      1762 ns/op
BenchmarkLogIoDiscardWriter-8                	 3000000	       443 ns/op
BenchmarkLogIoDiscardWriterWithoutFlags-8    	 5000000	       305 ns/op
BenchmarkLogCustomNullWriter-8               	 3000000	       436 ns/op
BenchmarkLogCustomNullWriterWithoutFlags-8   	 5000000	       301 ns/op
BenchmarkNopLogger-8                         	 500000	          3793 ns/op
BenchmarkNopLoggerWithoutFlags-8             	 1000000	      3859 ns/op
PASS
ok  	github.com/rafaelhenrique/golang_experiments/my_experiments/log	14.731s
```