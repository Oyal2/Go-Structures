# Go Structures <img src="https://user-images.githubusercontent.com/13637813/193483325-7b8b64c4-577d-43e1-a47c-ee0b75eb5bd0.png" width=300px align="right" >

Golang Data structures

## Benchmarks

emirpasic/gods ArrayList:
```
goos: windows
goarch: amd64
pkg: github.com/oyal2/Go-Structures/benchmark/GODS
cpu: AMD Ryzen 9 5950X 16-Core Processor            
BenchmarkInsert
BenchmarkInsert-32    	     100	  40914992 ns/op
BenchmarkGet
BenchmarkGet-32       	1000000000	         0.9050 ns/op
```

oyal2/GoStructure ArrayList:
```
goos: windows
goarch: amd64
pkg: github.com/oyal2/Go-Structures/benchmark/Go_Structrures
cpu: AMD Ryzen 9 5950X 16-Core Processor            
BenchmarkInsert
BenchmarkInsert-32    	    2499	    405562 ns/op
BenchmarkGet
BenchmarkGet-32       	1000000000	         0.9115 ns/op
```
