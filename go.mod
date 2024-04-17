module SDAS

go 1.22.1

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/vmihailenco/msgpack/v5 v5.4.1
	golang.org/x/image v0.0.0-20220302094943-723b81ca9867
)

require github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/stretchr/testify v1.9.0 // indirect
)
