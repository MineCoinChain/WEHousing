# PostOrders Service

This is the PostOrders service

Generated with

```
micro new IHome/PostOrders --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.PostOrders
- Type: srv
- Alias: PostOrders

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./PostOrders-srv
```

Build a docker image
```
make docker
```