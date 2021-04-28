EXE = shell
SRC = .
BIN = bin
RS_ADDR ?= 192.168.56.1:9092
LDFLAGS = -ldflags="-s -w -X main.Addr=$(RS_ADDR)"

all: build-windows build-darwin build-linux

bin:
	mkdir -p $(BIN)

build-windows: bin
	GOOS=windows go build -o $(BIN)/$(EXE)_windows.exe $(LDFLAGS) $(SRC)

build-%: bin
	GOOS=$* go build -o $(BIN)/$(EXE)_$* $(LDFLAGS) $(SRC)

clean:
	rm -rf $(BIN)