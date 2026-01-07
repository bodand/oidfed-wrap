CC ?= cc
CFLAGS ?= -Wall -Wextra -g

GOC := go build
GOFLAGS := -buildmode=c-archive -v
# Horrid, but not like Go cares about the best-practices
# of this makefile
GO_SRC := $(wildcard *.go)

PREFIX ?= /usr/local
INCLUDE_INSTALL ?= ${PREFIX}/include
LIBRARY_INSTALL ?= ${PREFIX}/lib

LIB_HEADER := oidfed_wrap_lib.h 
LIB_ARCHIVE := liboidfed_wrap.a

C_SRC := demo.c
C_OBJ := ${C_SRC:.c=.o}

.PHONY: build
build: ${LIB_HEADER} ${LIB_ARCHIVE}

demo: ${C_OBJ} ${LIB_ARCHIVE}
	${CC} -g ${CFLAGS} ${C_OBJ} -o $@ -L. -loidfed_wrap

${LIB_ARCHIVE}: oidfed_wrap_lib.a 
	cp oidfed_wrap_lib.a $@

oidfed_wrap_lib.h oidfed_wrap_lib.a: ${GO_SRC} oidfed_wrap.h
	${GOC} ${GOFLAGS} ${GO_SRC}

.PHONY: clean
clean:
	-rm -f ${LIB_ARCHIVE} ${LIB_HEADER} oidfed_wrap_lib.a
	-rm -f demo

.PHONY: install
install: build
	install -D -m644 -t ${INCLUDE_INSTALL} ${LIB_HEADER} 
	install -D -m644 -t ${LIBRARY_INSTALL} ${LIB_ARCHIVE}
