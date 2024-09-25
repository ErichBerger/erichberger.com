TDIR := ./ui/templates
UI := $(wildcard $(TDIR)/*.go)
WEB := $(wildcard ./cmd/web/*.go)

DEPS := $(UI) $(WEB)
APP := ./tmp/main

all: $(APP)

$(APP): $(DEPS) 
	go build --o ./tmp/main ./cmd/web

$(TDIR)/%_templ.go: $(TDIR)/%.templ
	templ generate