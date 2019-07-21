# GetUserOrder Service

This is the GetUserOrder service

Generated with

```
micro new IHome/GetUserOrder --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.GetUserOrder
- Type: srv
- Alias: GetUserOrder

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
./GetUserOrder-srv
```

Build a docker image
```
make docker
```