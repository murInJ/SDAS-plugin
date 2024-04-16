#!/bin/bash

#find . -maxdepth 2 -type f -name 'build.sh' -exec bash {} \;

# 如果你不想使用bash来执行脚本，可以使用sh：
 find . -maxdepth 2 -type f -name 'build.sh' -exec sh {} \;