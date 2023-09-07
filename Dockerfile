# とりあえずgolangのバージョンは最新を使用
FROM golang:latest
# アップデートとgitのインストール
RUN apt-get update && apt-get install -y git
RUN apt install sqlite3
WORKDIR /go/src

#go library
RUN go install github.com/99designs/gqlgen@latest
RUN go mod init github.com/hayabusa1228/graph_ql
RUN go get -u github.com/99designs/gqlgen
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest

#sqlite テーブル作成
RUN ./setup.sh

