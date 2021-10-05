# MS-API

## 準備

```
# GO111MODULEをオンにする
$ export GO111MODULE=on
```

## 動作

```
# docsを更新
$ swag init

# 実行
$ go run main.go
```

```
# ビルド
$ make build

# 実行
$ make run

# Dockerビルド
$ make docker-build

# Docker実行
$ make docker-run
```
