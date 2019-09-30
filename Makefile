build-local:
	sh ./hack/build.sh
generate-spec:
	quicktype ./spec.json -l go --package spec -o pkg/spec/spec.go -s schema -t CLISpec
example:
	rm -rf ../greet
	sh ./hack/example.sh
