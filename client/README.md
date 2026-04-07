# client

Vue ベースのフロントエンドです。

## できること

- `openapi-fetch` で型安全に API を呼び出せる
- OpenAPI から型生成して実装と仕様を同期できる
- `oxfmt` / `oxlint` / `vue-tsc` / build で品質チェックできる

## セットアップ

```bash
pnpm --dir client install
```

## 開発

```bash
mise run client
```

デフォルトの API ベース URL は `http://localhost:8080` です。
必要なら `VITE_API_BASE` を設定してください。

## OpenAPI 型生成

```bash
pnpm --dir client run codegen
```

生成ファイル:

- `src/gen/api-types.ts`

## Format / Lint / TypeCheck / Build

```bash
pnpm --dir client run fmt
pnpm --dir client run fmt:check
pnpm --dir client run lint
pnpm --dir client exec vue-tsc --noEmit
pnpm --dir client exec vp build
```

## 主なディレクトリ

- `src/lib/api.ts`: API クライアント
- `src/views`: 画面
- `src/components`: UI コンポーネント
- `src/gen`: OpenAPI 生成物
