SRCS := $(wildcard *.go)
PROGS := $(subst .go,,$(SRCS))

all:
	for src in $(SRCS); do go build $$src; done

clean:
	for prog in $(PROGS); do rm -f $$prog; done

.PHONY: all clean
