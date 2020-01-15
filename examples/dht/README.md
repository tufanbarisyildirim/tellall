## Distributed Hash Table
With distributed consistent hashing you always get same target with same key.

```shell script 
go run ./examples/dht
```

### Output of the example
```shell script
got message 0 to ch1
got message 1 to ch1
got message 2 to ch1
got message 3 to ch3
got message 4 to ch2
got message 5 to ch2
got message 6 to ch3
got message 7 to ch1
got message 8 to ch1
got message 9 to ch3
ch1 total message  5
ch2 total message  2
ch3 total message  3
```