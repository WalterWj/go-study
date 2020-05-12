---
title: Go modules
time: 2020年04月21日
auto: 王军
---
# 使用

## 创建 gomod

```shell
go mod init github.com/WalterWj/go-study
```

* `go run`, `go build`, `go test` 时，会自动下载相关依赖包。

## 好处

* gomod 可以语义话版本号，可以 git 对应分支或者 tag 的依赖包

```shell
go get github.com/go-sql-driver/mysql@v1.5.0
```

## 常用命令

* 查看所以依赖版本

```shell
$ go list -u -m all
github.com/WalterWj/go-study
github.com/go-sql-driver/mysql v1.5.0
```

* 升级

```shell
go get -u github.com/go-sql-driver/mysql  
```

* 只升级补丁

```shell
go get -u=patch github.com/go-sql-driver/mysql  
```

* 升降级版本

```shell

```
