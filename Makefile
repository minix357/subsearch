
APPNAME ?= subsearch
INSTDIR ?= /usr/local/bin

clean:
	if [ -f $(INSTDIR)/$(APPNAME) ]; then \
		rm $(INSTDIR)/$(APPNAME); \
	fi

build:
	go build -o $(APPNAME) .

install:
	mv $(APPNAME) $(INSTDIR)

all: clean build install
