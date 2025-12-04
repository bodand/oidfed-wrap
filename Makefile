GO_SRC := $(wildcard *.go)

demo: demo.c oidfed_wrap_lib.h liboidfed_wrap.a
	cc -g demo.c -Wall -Wextra -o $@ -L. -loidfed_wrap

oidfed_wrap_lib.h liboidfed_wrap.a: ${GO_SRC} oidfed_wrap.h
	go build -buildmode=c-archive -v ${GO_SRC}
	mv oidfed_wrap_lib.a liboidfed_wrap.a

