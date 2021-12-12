# MS-API

## 開発準備

```
# ffmpegをインストール
$ brew install ffmpeg

# GO111MODULEをオンにする
$ export GO111MODULE=on

# mod.modとgo.sumの差でエラーが出たら以下のコマンドを実行
$ go mod tidy

# docker用の共有ディレクトリを生成
$ mkdir -p ~/ms-tv/assets
$ mkdir -p ~/ms-tv/logs
```

## 実行コマンド

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
$ docker-compose up

# 停止コマンド
$ docker-compose down
```

## ドキュメント

http://localhost:8080/docs/api/v1/index.html#/
