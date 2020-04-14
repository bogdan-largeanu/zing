args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`
# Testing functions for cobra https://www.bradcypert.com/testing-a-cobra-cli-in-go/
default: demo

demo:
	go run zing.go demo

#type arguments after profile
run:
	go run zing.go $(call args, ;)

test:
	go test -v