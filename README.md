# Tell-all
Tellall (Bellman) is a golang channel library helps you to;

- [x] easily implement pub/sub mechanism
- [x] distribute messages between channels using consistent hash algorithm
- [ ] balance messages between channels using some [balancing algorithms](https://kemptechnologies.com/load-balancer/load-balancing-algorithms-techniques/)
    - [x] round-robin
    - [x] weighted round robin
    - [ ] agent-based adaptive load balancing

### Examples

- [Pub/Sub](examples/pubsub/README.md)
- [Balancing](examples/balance/README.md)
- [Consistent Hashing](examples/dht/README.md)
