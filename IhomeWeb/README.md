# IhomeWeb Service

This is the IhomeWeb service

Generated with

```
micro new sss/IhomeWeb --namespace=go.micro --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.web.IhomeWeb
- Type: web
- Alias: IhomeWeb

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
./IhomeWeb-web
```

Build a docker image
```
make docker
```