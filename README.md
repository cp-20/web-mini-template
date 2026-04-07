# web-mini-template

春ハッカソン向けの Web アプリ開発テンプレートです。

## このテンプレートでできること

- Vue + Go + MySQL の構成をすぐに立ち上げられる
- OpenAPI を単一の仕様として client/server コード生成できる
- 型安全な API 呼び出し（`openapi-fetch`）でフロント実装できる
- `oxfmt` / `oxlint` / TypeCheck / Build を CI で分割チェックできる
- server 側を testcontainers の e2e で検証できる

## まずやること

```bash
mise trust
mise install
pnpm --dir client install
```

## よく使うコマンド

- まとめて開発起動（client + server）

```bash
mise run dev
```

- 停止

```bash
mise run dev-down
```

- OpenAPI からコード生成

```bash
mise run codegen
```

- client の整形・lint

```bash
mise run client-fmt
mise run client-fmt-check
mise run client-lint
```

- server の build/test

```bash
mise run server-build
mise run server-test
```

## 詳細ドキュメント

- client 開発手順: `client/README.md`
- server 開発手順: `server/README.md`

## 構成

- `client`: フロントエンド
- `server`: バックエンド
- `openapi`: API 仕様
- `.github/workflows/ci.yml`: CI
