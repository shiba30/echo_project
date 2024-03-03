# Twitter(X) クローン

## 使用技術一覧

<!-- シールド一覧 -->
  <!-- フロントエンド一覧 -->
  <img src="https://img.shields.io/badge/-Next.js-000000.svg?logo=next.js&style=plastic">
  <img src="https://img.shields.io/badge/-React-61DAFB.svg?logo=react&style=plastic">
  <img src="https://img.shields.io/badge/-Typescript-007ACC.svg?logo=typescript&style=plastic">
  <img src="https://img.shields.io/badge/tailwindcss-0F172A?&logo=tailwindcss">
  <!-- バックエンド一覧 -->
  <img src="https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic">
  <!-- ミドルウェア一覧 -->
  <img src="https://img.shields.io/badge/-Postgresql-336791.svg?logo=postgresql&style=popout">
  <!-- インフラ一覧 -->
  <img src="https://img.shields.io/badge/-Docker-1488C6.svg?logo=docker&style=popout">

## 目次

1. [システム概要](#システムについて)
2. [設計](#設計)
3. [環境](#環境)
4. [開発環境構築](#開発環境構築)

## システム概要

#### Twitter(X) Clone

- スクール課題
- フロントエンドの Next はこちら（[Next.js](https://github.com/shiba30/twitter_nextjs_frontend)）
- バックエンドの Echo はクリーンアーキテクチャで作成します

## 設計

- 以前、Gin のみで Twitter クローン（[リポジトリ](https://github.com/shiba30/twitter_clone_gin) ）を作成した際、機能追加時に依存関係のロジック関係で修正が大変だったので、なるべく将来性を見越して改修ボリュームを最小限に抑えることができる設計を取り入れることにした。
  - Clean Architecture を採用
  - UI 設計は Atomic Design

## 環境

    | 言語・フレームワーク | バージョン |
    | -------------------- | ---------- |
    | Next.js              | 14.1.0     |
    | React                | 18.2.0     |
    | typescript           | 5.3.3      |
    | tailwindcss          | 3.4.1      |
    | Go                   | 1.20       |
    | Echo                 | 4.11.4     |
    | Postgresql           | 15         |

その他のパッケージのバージョンは package.json を参照してください

## 開発環境構築

1. 本レポジトリを任意のディレクトリに `$ git clone` する
2. プロジェクトルート配下で、以下の依存関係のインストールを行う

   ```
   $ npm install
   ```

3. Node.js のインストール
   - [Node.js](https://nodejs.org/en)
   - 推奨版（安定版）をインストール
4. .env ファイルの環境変数定義ファイルに以下の DB 接続情報を定義して作成する

   ```
   DB_USER="ユーザ名"
   DB_PASSWORD="パスワード"
   DB_NAME="データベース名"
   DB_PORT="ポート番号"
   ```

### 開発環境実行

- 起動

  - ルートディレクトリ配下で実行

    ```
    $ docker-compose up
    ```

### 動作確認

http://localhost:1323/ にブラウザからアクセスできたら成功
