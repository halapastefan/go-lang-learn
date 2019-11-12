* Install Protocol Buffers v3

Next, install the protoc plugin for Go

$ go get -u github.com/golang/protobuf/protoc-gen-go

The compiler plugin, protoc-gen-go, will be installed in $GOBIN, defaulting to $GOPATH/bin. It must be in your $PATH for the protocol compiler, protoc, to find it.