NAME=infoscan
BINDIR=bin
TARGET=target
VERSION=$(shell git describe --tags --always || echo "unknown version")
BUILDTIME=$(shell date -u)
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-X "github.com/Ymjie/GScan/constant.Version=$(VERSION)" \
		-X "github.com/Ymjie/GScan/constant.BuildTime=$(BUILDTIME)" \
		-w -s -buildid='

PLATFORM_LIST = \
	darwin-amd64 \
	darwin-arm64 \
	linux-amd64 \

WINDOWS_ARCH_LIST = \
	windows-amd64 \

all: linux-amd64 darwin-amd64 windows-amd64 # Most used

docker:
	$(GOBUILD) -o $(BINDIR)/$(NAME)-$@

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$@/$(NAME)-$@

darwin-arm64:
	GOARCH=arm64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$@/$(NAME)-$@

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$@/$(NAME)-$@

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$@/$(NAME)-$@.exe


gz_releases=$(addsuffix .gz, $(PLATFORM_LIST))
zip_releases=$(addsuffix .zip, $(WINDOWS_ARCH_LIST))

$(gz_releases): %.gz : %
	chmod +x $(BINDIR)/$(basename $@)/$(NAME)-$(basename $@)
	cp ./config.yml keywords.txt url.txt whitelist.txt $(BINDIR)/$(basename $@)/
	tar -zcvf $(TARGET)/$(NAME)-$(basename $@)-$(VERSION).gz ./$(BINDIR)/$(basename $@)/

$(zip_releases): %.zip : %
	cp ./config.yml keywords.txt url.txt whitelist.txt $(BINDIR)/$(basename $@)/
	zip -m -r -j $(TARGET)/$(NAME)-$(basename $@)-$(VERSION).zip $(BINDIR)/$(basename $@)/
	
all-arch: $(PLATFORM_LIST) $(WINDOWS_ARCH_LIST)

releases: $(gz_releases) $(zip_releases)

