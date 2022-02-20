# MS-API

## 開発準備

```
# GO111MODULEをオンにする
$ export GO111MODULE=on

# App初期設定
$ make init

# ffmpegをインストール
$ brew install ffmpeg

# mod.modとgo.sumの差でエラーが出たら以下のコマンドを実行
$ go mod tidy
```

## 実行コマンド

```
# docsを更新
$ swag init

# 実行
$ make run

# Dockerビルド
$ make docker-build

# Docker実行
$ make docker-run

# external appsを立ち上げる
$ make external-run

# external appsを終了する
$ make external-end

# ヘルプ
$ make help
```

## ドキュメント

http://localhost:8080/docs/api/v1/index.html#/
