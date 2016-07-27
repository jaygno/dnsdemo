all: test 

uninstall: remove

test: test.go mac.go parse.go dnsmsg.go
	go build -o test test.go mac.go parse.go dnsmsg.go 

clean:
	rm -f ./test
