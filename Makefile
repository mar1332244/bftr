.PHONY: all
all:
	go build -v -o /usr/local/bin/bftr main.go

.PHONY: clean
clean:
	rm /usr/local/bin/bftr
