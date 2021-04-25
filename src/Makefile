args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`
# Testing functions for cobra https://www.bradcypert.com/testing-a-cobra-cli-in-go/
default: demo

demo:
	go run zing.go demo s
tele:
	go run zing.go teleport s

#type arguments after profile
run:
	go run zing.go $(call args, ;)

test:
	go test -v

install: build clean
	@sudo ln -s /Users/bogdanlargeanu/go/src/zing/src/zing /usr/local/bin

build:
	@echo "Building Zing"
	@go build zing.go

clean:
	@sudo rm /usr/local/bin/zing
