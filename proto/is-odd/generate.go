package is_odd

//go:generate sh -c "protoc -I $(go list -m -f '{{.Dir}}' github.com/q3k/is-odd)/proto/is-odd --go_out=plugins=grpc:. is-odd.proto"
