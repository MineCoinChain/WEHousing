# PutOrders Service

This is the PutOrders service

Generated with

```
micro new IHome/PutOrders --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.PutOrders
- Type: srv
- Alias: PutOrders

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
./PutOrders-srv
```

Build a docker image
```
make docker
```