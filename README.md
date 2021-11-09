# MS-API

## 準備

```
# GO111MODULEをオンにする
$ export GO111MODULE=on

# mod.modとgo.sumの差でエラーが出たら以下のコマンドを実行
$ go mod tidy
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

# Docker Compose系
$ make docker-compose-build
$ make docker-compose-up
$ make docker-compose-down
$ make docker-compose-run // build & up
```

## ドキュメント

http://localhost:8080/docs/api/v1/index.html#/
