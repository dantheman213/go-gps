.PHONY: deps

deps:
	@echo Downloading go.mod dependencies && \
		go mod download
