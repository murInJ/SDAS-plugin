go build -ldflags="-s -w" --buildmode=plugin -o ./plugins/VideoWall.so ./VideoWall/plugin_VideoWall.go
chmod +x ./plugins/VideoWall.so
upx -9 ./plugins/VideoWall.so
chmod -x ./plugins/VideoWall.so
