# 鲸管家 WhaMan

## Introduction

进销存管理系统

## Getting Started

1. 准备相关环境和工具：

- Go 1.16
- MySQL 8.0

2. 获取ssl证书，可使用Let's Encrypt免费证书

3. 修改配置文件`./conf.yml`

4. 生成 swagger 文档：

```
// go get -u github.com/swaggo/swag/cmd/swag
swag init
```

5. 运行

```
go run main.go
```

or

```
go build WhaMan
```
