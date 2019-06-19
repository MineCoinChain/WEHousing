# PostUserAuth Service

This is the PostUserAuth service

Generated with

```
micro new sss/PostUserAuth --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.PostUserAuth
- Type: srv
- Alias: PostUserAuth

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
./PostUserAuth-srv
```

Build a docker image
```
make docker
```