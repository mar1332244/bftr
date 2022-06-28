SRC=./src
TARGET=/usr/local/bin/bftr

.PHONY: all
all:
	go build -o $(TARGET) $(SRC)

.PHONY: clean
clean:
	rm /usr/local/bin/bftr
