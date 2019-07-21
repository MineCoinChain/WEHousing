# GetHouseInfo Service

This is the GetHouseInfo service

Generated with

```
micro new IHome/GetHouseInfo --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.GetHouseInfo
- Type: srv
- Alias: GetHouseInfo

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
./GetHouseInfo-srv
```

Build a docker image
```
make docker
```