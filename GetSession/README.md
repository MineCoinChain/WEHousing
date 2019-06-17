# GetSession Service

This is the GetSession service

Generated with

```
micro new sss/GetSession --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.GetSession
- Type: srv
- Alias: GetSession

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
./GetSession-srv
```

Build a docker image
```
make docker
```