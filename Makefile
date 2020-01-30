args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

default: test

test:
	cd src && \
	go run main.go

#type arguments after profile
run:
	cd src && \
	go run main.go $(call args, ;)
