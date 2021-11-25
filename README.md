# MS-API

## 準備

```
# GO111MODULEをオンにする
$ export GO111MODULE=on

# mod.modとgo.sumの差でエラーが出たら以下のコマンドを実行
$ go mod tidy
```

## ビルド

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

## ms-tvを実行

```
# ms-api, ms-stream-api, ms-recommendation-apiのそれぞれで make docker-build を実行してから、以下を実施してください
$ make docker-compose-up

# 停止コマンド
$ make docker-compose-down
```

## ドキュメント

http://localhost:8080/docs/api/v1/index.html#/
