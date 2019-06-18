# GetUserInfo Service

This is the GetUserInfo service

Generated with

```
micro new sss/GetUserInfo --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.GetUserInfo
- Type: srv
- Alias: GetUserInfo

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
./GetUserInfo-srv
```

Build a docker image
```
make docker
```