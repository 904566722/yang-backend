compile-yang:
	CGO_ENABLED=0 GO111MODULE=on go build -mod vendor -o _out/bin/yang cmd/yang.go