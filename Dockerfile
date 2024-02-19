####################### Development stage #######################
FROM golang:1.20.5-alpine3.18 AS development

# 作業ディレクトリの定義
WORKDIR /app

# go.mod と go.sum を app ディレクトリにコピー
COPY go.mod go.sum ./

# 指定されたモジュールをダウンロード
RUN go mod download

# ルートディレクトリの中身を app フォルダにコピー
COPY . .

# ビルドするファイルを指定：main.go
RUN go build -o main /app/main.go

# air インストール
RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air

# .air.toml ファイルをコピー
COPY .air.toml ./

CMD ["air", "-c", ".air.toml"]

####################### Production stage #######################
FROM alpine:3.18 AS production

# 作業ディレクトリの定義
WORKDIR /app

# Development stage からビルドされた main だけを Production stage にコピー
COPY --from=development /app/main .

CMD ["./main"]
