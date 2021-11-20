# The Minimal Blockchain

In progress.

## Feature

- No third-party modules
- A little unstable philosophy
- Source code valued itself rather than engeering
- The consensus implements [*A consensus mechanism based on "egocentrism"*](https://smallyu-net.translate.goog/2021/10/29/%E4%B8%80%E7%A7%8D%E5%9F%BA%E4%BA%8E%E2%80%9C%E8%87%AA%E6%88%91%E4%B8%AD%E5%BF%83%E4%B8%BB%E4%B9%89%E2%80%9D%E7%9A%84%E5%85%B1%E8%AF%86%E6%9C%BA%E5%88%B6/?_x_tr_sch=http&_x_tr_sl=auto&_x_tr_tl=en&_x_tr_hl=en-US&_x_tr_pto=nui)

## Quick start

Start single node in local:

```
go run .
```

Start a three node group for test:

```

```

## Configuration

You change the config by modifying the const in the "config/config.go" or setup environment variable.

Support environment variable list:

```

```


## HTTP API

|Path|HTTP Method|Description|
|--|--|--|
|/post|POST|Post data to node|

## API Example

### /post
```
curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"key": "value"}' \
    http://127.0.0.1:25001/post
```
