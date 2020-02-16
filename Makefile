.POSIX:
.SUFFIXES:
.SUFFIXES: .1 .5 .7 .1.scd .5.scd .7.scd

VERSION=0.1.0

VPATH=doc
PREFIX?=/usr/local
_INSTDIR=$(DESTDIR)$(PREFIX)
BINDIR?=$(_INSTDIR)/bin
SHAREDIR?=$(_INSTDIR)/share/te
MANDIR?=$(_INSTDIR)/share/man
GO?=go
GOFLAGS?=

GOSRC!=find . -name '*.go'
GOSRC+=go.mod go.sum

te: $(GOSRC)
	$(GO) build $(GOFLAGS) \
		-ldflags "-X main.Prefix=$(PREFIX) \
		-X main.ShareDir=$(SHAREDIR) \
		-X main.Version=$(VERSION)" \
		-o $@

te.conf: config/te.conf.in
	sed -e 's:@SHAREDIR@:$(SHAREDIR):g' > $@ < config/te.conf.in

DOCS := \
	te.1

.1.scd.1:
	scdoc < $< > $@

doc: $(DOCS)

all: te te.conf doc

# Exists in GNUMake but not in NetBSD make and others.
RM?=rm -f

clean:
	$(RM) $(DOCS) te.conf te

install: all
	mkdir -m755 -p $(BINDIR) $(MANDIR)/man1 \
		$(SHAREDIR)
	install -m755 te $(BINDIR)/te
	install -m644 te.1 $(MANDIR)/man1/te.1
	install -m644 te.conf $(SHAREDIR)/te.conf

RMDIR_IF_EMPTY:=sh -c '\
if test -d $$0 && ! ls -1qA $$0 | grep -q . ; then \
	rmdir $$0; \
fi'

uninstall:
	$(RM) $(BINDIR)/te
	$(RM) $(MANDIR)/man1/te.1
	$(RM) -r $(SHAREDIR)
	${RMDIR_IF_EMPTY} $(BINDIR)
	$(RMDIR_IF_EMPTY) $(MANDIR)/man1
	$(RMDIR_IF_EMPTY) $(MANDIR)

.DEFAULT_GOAL := all

.PHONY: all doc clean install uninstall
