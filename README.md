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

Reference the documents part in [the-minimal-blockchain/tmb-specification](https://github.com/the-minimal-blockchain/tmb-specification#Documents).
