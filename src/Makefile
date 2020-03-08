args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`
# Testing functions for cobra https://www.bradcypert.com/testing-a-cobra-cli-in-go/
default: demo

demo:
	go run main.go demo

#type arguments after profile
run:
	go run main.go $(call args, ;)

test:
	go test -v