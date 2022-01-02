# The Minimal Blockchain

It exactly is a blockchain.

## Feature

- No third-party modules
- A little unstable philosophy
- Source code valued itself rather than engeering
- The consensus implements [*A consensus mechanism based on "egocentrism"*](https://smallyu-net.translate.goog/2021/10/29/%E4%B8%80%E7%A7%8D%E5%9F%BA%E4%BA%8E%E2%80%9C%E8%87%AA%E6%88%91%E4%B8%AD%E5%BF%83%E4%B8%BB%E4%B9%89%E2%80%9D%E7%9A%84%E5%85%B1%E8%AF%86%E6%9C%BA%E5%88%B6/?_x_tr_sch=http&_x_tr_sl=auto&_x_tr_tl=en&_x_tr_hl=en-US&_x_tr_pto=nui)

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

## Configuration

You change the config by modifying the const in the `config/config.go` or setup environment variable.

Supported environment variable list:

| Variable name | Description |
| -- | -- |
| LocalPort | Node local port |
| DataPath | Block data path |
| HTTPPort | Node http port |

## HTTP API

| Path | HTTP Method | Description |
| -- | -- | -- |
| /info | GET | Get node info |
| /post | POST | Post data to node |
| /get/{blockHeigth} | GET | Get block data from node |

## API Example

### /info

```
curl http://127.0.0.1:25010/info
```

### /post

```
curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"key": "value"}' \
    http://127.0.0.1:25010/post
```

### /get

```
curl http://127.0.0.1:25010/get?height=1
```
