# SDAS-plugin
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