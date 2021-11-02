#!/usr/bin/env bash

def_version="1.0.0"

version=''
os=''
arch=''
debug=''

echo -e "请选择操作系统:"
oss=("linux" "windows" "darwin")
select opt in "${oss[@]}"; do
  case $opt in
  "linux")
    os=$opt
    break
    ;;
  "windows")
    os=$opt
    break
    ;;
  "darwin")
    os=$opt
    break
    ;;
  *) echo "invalid option $REPLY" ;;
  esac
done

echo -e "请选择CPU架构:"
archs=("amd64" "arm" "arm64")
select opt in "${archs[@]}"; do
  case $opt in
  "amd64")
    arch=$opt
    break
    ;;
  "arm")
    arch=$opt
    break
    ;;
  "arm64")
    arch=$opt
    break
    ;;
  *) echo "invalid option $REPLY" ;;
  esac
done

read -p "请选择编译版本:" name

echo -e "是否是调试版本"
debugs=("false" "true")
select opt in "${debugs[@]}"; do
  case $opt in
  "true")
    debug=''
    break
    ;;
  "false")
    debug="-ldflags '-s -w'"
    break
    ;;
  *) echo "invalid option $REPLY" ;;
  esac
done

export CGO_ENABLED=1
export GOOS=${os}
export GOARCH=${arch}

# CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w'
build_info="CGO_ENABLED=1 GOOS=${os} GOARCH=${arch} go build ${debug}"
echo "编译指令为:${build_info}"

build_cmd="go build ${debug}"

`${build_cmd}`

