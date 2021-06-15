# SolidFire SDK

SDK for working with the NetApp SolidFire Element OS (still v0)

Solidfire Element OS API Doc: https://docs.netapp.com/us-en/element-software/


### Running Tests

Use `make test` to run all tests. This will always run unit tests.

To run integration tests, you need to define the following environment variables:

```
SOLIDFIRE_HOST
SOLIDFIRE_HOST2
SOLIDFIRE_PASS
SOLIDFIRE_USER
```

Otherwise the integration tests will be skipped. The `SOLIDFIRE_HOST` and `SOLIDFIRE_HOST2` values should be set to the MVIP of two different test clusters.

### Client examples

See /examples/main.go for example client code that instantiates and uses this SDK.
