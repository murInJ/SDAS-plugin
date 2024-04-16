# SDAS-plugin
plugin hub of [SDAS](https://github.com/murInJ/Streaming-Data-Aggregation-Source-Service)

Notice:

如果你想开发新的插件

请将接口函数命名为`Pipeline(entry map[string]*api.SourceMsg) (map[string]*api.SourceMsg,error)`

同时留有`Init(args string) error`函数

plugin list:
- Video Wall