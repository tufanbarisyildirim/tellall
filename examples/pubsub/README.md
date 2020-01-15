## Channel Balancing
Pub/Sub

```shell script 
go run ./examples/pubsub
```

### Output of the example
```shell script
Sub2[b975a237] got message  1: 0 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  1: 0 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  2: 0 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  2: 0 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  1: 1 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  1: 1 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  2: 1 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  2: 1 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  1: 2 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  1: 2 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  2: 2 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  2: 2 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  2: 3 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  2: 3 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  1: 3 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  1: 3 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  2: 4 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  2: 4 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  2: 5 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  2: 5 from Pub[183eeb4d23e3e34a]
```