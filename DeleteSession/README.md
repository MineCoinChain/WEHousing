# DeleteSession Service

This is the DeleteSession service

Generated with

```
micro new sss/DeleteSession --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.DeleteSession
- Type: srv
- Alias: DeleteSession

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
./DeleteSession-srv
```

Build a docker image
```
make docker
```