.PHONY: all
all:
	go build -o /usr/local/bin/bftr main.go

.PHONY: clean
clean:
	rm /usr/local/bin/bftr
