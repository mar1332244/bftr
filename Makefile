.PHONY: all
all:
	go build -o bftr main.go

.PHONY: clean
clean:
	rm /usr/local/bin/bftr
