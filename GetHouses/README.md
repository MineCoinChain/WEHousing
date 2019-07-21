# GetHouses Service

This is the GetHouses service

Generated with

```
micro new IHome/GetHouses --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.GetHouses
- Type: srv
- Alias: GetHouses

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
./GetHouses-srv
```

Build a docker image
```
make docker
```