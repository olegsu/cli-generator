.PHONY: build
build:
	sh ./scripts/build.sh

.PHONY: generate-spec
generate-spec:
	quicktype ./spec.json -l go --package spec -o pkg/spec/spec.go -s schema -t CLISpec

example:
	rm -rf ./greet || true
	sh ./scripts/example.sh
example-2:
	rm -rf ./cli-example-sub-command || true
	sh ./scripts/cli-example-sub-command.sh