PingFederate SDK Go
==================

- ![Status: Experimental](https://img.shields.io/badge/status-experimental-EAAA32) ![ci](https://github.com/iwarapter/pingfederate-sdk-go/workflows/ci/badge.svg)

This project was created to support the [terraform-provider-pingfederate](https://github.com/iwarapter/terraform-provider-pingfederate) and other experimental projects.

The SDK is currently generated, however it is extremely experimental and subject to change without notice.

Using the SDK
----------------------

```
go get github.com/iwarapter/pingfederate-sdk-go@master
```

Testing the SDK
---------------------------

In order to test the provider, you can run `make test`.

```sh
$ make test
```

This will run the acceptance tests by initializing a local docker container to execute the functional tests against.
