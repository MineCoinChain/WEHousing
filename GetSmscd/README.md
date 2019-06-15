# GetSmscd Service

This is the GetSmscd service

Generated with

```
micro new sss/GetSmscd --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.GetSmscd
- Type: srv
- Alias: GetSmscd

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
./GetSmscd-srv
```

Build a docker image
```
make docker
```