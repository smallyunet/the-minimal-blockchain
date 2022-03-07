# The Minimal Blockchain

The Golang implementation of [the-minimal-blockchain/tmb-specification](https://github.com/the-minimal-blockchain/tmb-specification).

## Environment

- go version go1.17 darwin/arm64

## Quick start

Start single node in local:

```
go run .
```

Start a 3 node's group in 3 terminal windows:

```
export LocalPort=25000 DataPath=data1 HTTPPort=25010
go run .

export LocalPort=25001 DataPath=data2 HTTPPort=25011
go run .

export LocalPort=25002 DataPath=data3 HTTPPort=25012
go run .
```

## Documents

### Configuration

You can change the config by modifying the const in the config file or setting up the environment variable.

Supported environment variable list currently:

| Variable name | Description |
| -- | -- |
| LocalPort | Node local port |
| LocalIP | Node local ip |
| DataPath | Block data path |
| HTTPPort | Node http port |

### HTTP API

| Path | HTTP Method | Description |
| -- | -- | -- |
| /info | GET | Get node info |
| /post | POST | Post data to node |
| /get/{blockHeigth} | GET | Get block data from node |

### HTTP API Example

#### /info

```
curl http://127.0.0.1:25010/info
```

#### /post

```
curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"key": "value"}' \
    http://127.0.0.1:25010/post
```

#### /get

```
curl http://127.0.0.1:25010/get/1
```

