# Go Structures <img src="https://user-images.githubusercontent.com/13637813/193483325-7b8b64c4-577d-43e1-a47c-ee0b75eb5bd0.png" width=300px align="right" >

Golang Data structures

## Benchmarks

Benchmark Specs 
```
goos: windows
goarch: amd64
cpu: AMD Ryzen 9 5950X 16-Core Processor     
```
### ArrayList
emirpasic/gods ArrayList:
```      
BenchmarkInsert   	       100	       41555013 ns/op
BenchmarkGet        1000000000	          0.9050 ns/op
BenchmarkUpdate    	 245901840	         4.935  ns/op
BenchmarkRemove    	1000000000	         0.9305 ns/op
```

GoStructure ArrayList:
```       
BenchmarkInsert
BenchmarkInsert-32    	    3433	    374162 ns/op
BenchmarkGet
BenchmarkGet-32       	1000000000	         0.9045 ns/op
BenchmarkUpdate
BenchmarkUpdate-32    	539326327	         2.123 ns/op
BenchmarkRemove
BenchmarkRemove-32    	 3640387	       324.0 ns/op
```
