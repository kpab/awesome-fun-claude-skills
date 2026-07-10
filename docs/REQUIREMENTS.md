# awesome-fun-claude-skills 要件

## 背景・目的
ネタ系・エンタメ系の Claude Code スキル(ペルソナ・ジョーク・遊び系)は個別記事としてバズるのに集約する場所がない。
汎用の awesome リストは実務特化でこの領域を扱っておらず、ニッチが空いている(詳細は docs/IDEA.md 参照)。
「ネタ系特化」で一番のキュレーションリストになることを狙う。英語ベースのグローバル向け。

**成功条件**: YAML ソースから生成された README(awesome リスト)が 20〜30 件のネタ系スキルを分類掲載して GitHub で公開されている。

## ユーザー
不特定多数に公開(グローバル、英語ベース)。日本語圏のキャラ系スキル文化を英語圏に紹介する立ち位置。

## スコープ

### やること(MVP)
- ネタ系スキルのキュレーション: GitHub・Zenn・Qiita 等に散在するスキル/記事を収集し、カテゴリ分類して掲載(初回 20〜30 件)
- YAML 単一ソース → README(awesome リスト形式)の自動生成

### やらないこと(今回は)
- 閲覧サイト(M3 以降。作るときは Astro + Cloudflare、一覧+カテゴリ+検索)
- 自作スキルの大量生産(キュレーションが主軸)
- 投稿フォーム・レビュー・ランキング等の動的機能
- plugin marketplace 形式での配布(将来 marketplace.json を足す拡張余地のみ意識)

## 技術選定

| 項目 | 選定 | 理由 |
|------|------|------|
| データ | YAML(`data/skills.yaml`) | README とサイトの単一ソース。二重管理を排除 |
| 生成スクリプト | Go | 主戦場スタック。単バイナリで CI でも動かしやすい |
| フロントエンド | なし(M3 で Astro 予定) | まず README で反応を見る |
| インフラ | GitHub のみ(M3 で Cloudflare 予定) | v1 は静的リポジトリで完結 |
| CI | GitHub Actions | YAML→README の生成差分チェック(生成漏れ防止) |

## 公開・運用
- リポジトリ: public(GitHub: kpab)/ ライセンス: リストコンテンツ CC0、生成スクリプト MIT
- 言語: README は英語メイン(日本語版 README.ja.md は将来、同じ YAML から生成)
- デプロイ先: なし(GitHub リポジトリ自体が成果物)
- CI: GitHub Actions — YAML の lint と README 生成が最新かのチェック

## マイルストーン

- **M1(最小の動くもの)**: YAML スキーマ確定+生成スクリプトが動き、手元の素材(調査で見つけた約10件+自作2件)を載せた README が生成される
  - [x] genreadme を実装する(yaml.v3 でパース → カテゴリ別に README 生成。category 参照切れ・URL 重複・必須欠落はエラー)
  - [x] `go run ./cmd/genreadme` で README.md が生成され、プレースホルダが置き換わることを確認する
  - [ ] /collect を1周回して掲載を 15 件程度まで増やす
- **M2(初回公開)**: 掲載 20〜30 件に拡充、カテゴリ整備、CONTRIBUTING.md、CI 整備、public 公開
- **M3(サイト)**: Astro + Cloudflare で閲覧サイト(一覧+カテゴリ+クライアントサイド検索)。同じ YAML から生成

## 将来(今は設計だけ意識)
- 自作ネタスキル(博多弁ギャル・知的ギャル等)の掲載(一旦不要と判断。載せる場合は公開方法を要検討)
- スキル詳細ページ(口調サンプル・インストールコマンド付き)— YAML に example/install フィールドの拡張余地を残す
- marketplace.json による /plugin 直接インストール対応
- 日本語版 README(README.ja.md)の併設

## 未定事項
- カテゴリ体系の確定(persona / joke / game / meme 等 — 収集しながら決める)
- リポジトリ名の最終確定(awesome-fun-claude-skills 仮。awesome-lint 準拠にするかも含め M2 で判断)
- 自作スキルの同梱方法(リポジトリ内に置くか、別リポジトリ参照か)
