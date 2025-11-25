GO_SRC := $(wildcard *.go)

demo: demo.c main.a main.h
	cc -g demo.c -Wall -Wextra -o $@ -L. -l:main.a

main.a main.h: ${GO_SRC} oidfed_wrap.h
	go build -buildmode=c-archive -v ${GO_SRC}

