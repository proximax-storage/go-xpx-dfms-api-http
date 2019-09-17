# go-xpx-dfms-http-client

[![Documentation](https://godoc.org/github.com/proximax-storage/go-xpx-dfms-http-client?status.svg)](https://godoc.org/github.com/proximax-storage/go-xpx-dfms-http-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/proximax-storage/go-xpx-dfms-http-client)](https://goreportcard.com/report/github.com/proximax-storage/go-xpx-dfms-http-client)
[![proximax](https://img.shields.io/badge/project-ProximaX-orange)](https://www.proximax.io/)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/proximax-storage/go-xpx-dfms-http-client)

> *The Client's API is experimental and can often change.*
>
> **Current status**: under active development 

Welcome to DFMS! 

The package is the HTTP client for all DFMS's applications. It gives an ability to 
operate DFMS(Clients) and DFMSR(Replicators) nodes remotely through a generic API.

## Table of Contents

- [Background](#background)
- [Requirements](#requirements)
- [Install](#install)
- [Usage](#usage)
- [API](#api)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Background

DFMS(Decentralized File Management System) - is a new advanced technology 
developed by [ProximaX](https://www.proximax.io/) which represents the Storage Layer of ProximaX infrastructure platform solution. 

DFMS provides an easy-to-use decentralized market of disk space using the Blockchain and the DLT
powdered with cryptography magic. 

## Requirements

    go version 1.12 or higher

## Install

`$ go get github.com/proximax-storage/go-xpx-dfms-http-client`

## Usage

Simply create new Client using an address the DFMS application's API listens to
and you are ready to go:

```go
client := client.New(address)

...
```

Creating new Drive contract:

```go
contract, err := client.ContractAPI().Compose(ctx, space, duration)

...
```

## API
<!---
Add link to an external API repository
-->

TODO

## Maintainers

[@Wondertan](https://github.com/Wondertan)
[@BramBear](https://github.com/alvin-reyes)
[@mrLSD](https://github.com/mrLSD)

## Contributing

Feel free to dive in! [Open an issue](https://github.com/proximax-storage/go-xpx-dfms-http-client/issues/new) or submit PRs.

## License

[Apache 2.0](LICENSE) © 2019 ProximaX Limited
