SHELL=/bin/zsh

actualdir=$(shell dirname `pwd`)

init:
	export GOPATH="$(actualdir)"

run:
	ENVIRONMENT=local go run .
