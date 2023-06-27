GO     := helper/usr/local/bin/go
HELPER := helper

$(HELPER)/lib/cache/docker-setup/metadata.json:
	@docker-setup --prefix=$(HELPER) update

$(HELPER)/lib/cache/docker-setup/manifests/%.json: \
		$(HELPER)/lib/cache/docker-setup/metadata.json
	@docker-setup --prefix=$(HELPER) install $*

demo: cmd/demo/*.go $(HELPER)/lib/cache/docker-setup/manifests/go.json
	@$(GO) build -o $@ ./cmd/demo
