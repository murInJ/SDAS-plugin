# SDAS-plugin
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/murInJ/SDAS-plugin.svg)](https://github.com/murInJ/go-pic-bed)
![GitHub Release](https://img.shields.io/github/v/release/murInJ/SDAS-plugin)
[![GitHub contributors](https://img.shields.io/github/contributors/MurInJ/SDAS-plugin.svg)](https://GitHub.com/MurInJ/SDAS-plugin/graphs/contributors/)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/murInJ/SDAS-plugin/go.yml)
plugin hub of [SDAS](https://github.com/murInJ/Streaming-Data-Aggregation-Source-Service)

Notice:

如果你想开发新的插件

请将接口函数命名为`Pipeline(entry map[string]any) (map[string]any,error)`
其中`map[string]any`定义为
```go
{
	Ntp int64
	Data []byte
	DataType string
}
```


同时留有`Init(args string) error`函数

plugin list:
- Video Wall