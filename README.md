# GraphQL + gRPC Monorepo

このプロジェクトは、GraphQLとgRPCを使用したマイクロサービスアーキテクチャのモノレポです。

## プロジェクト構成

```
.
├── frontend/          # Next.js + React フロントエンド
├── backend/           # Golang バックエンド
└── packages/
    └── proto/        # gRPC の共有プロトコル定義
```

## 開発環境のセットアップ

1. フロントエンド（Next.js）のセットアップ
```bash
cd frontend
npm install
```

2. バックエンド（Golang）のセットアップ
```bash
cd backend
go mod init
go mod tidy
```

## 開発サーバーの起動

- フロントエンド開発サーバー:
```bash
npm run dev:frontend
```

- バックエンド開発サーバー:
```bash
npm run dev:backend
```
