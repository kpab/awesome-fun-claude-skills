# awesome-fun-claude-skills

ネタ系・エンタメ系 Claude Code スキル(ペルソナ・ジョーク・遊び系)の awesome リスト。
YAML 単一ソースから README を自動生成する。英語メイン・グローバル向け。

要件・計画の詳細は docs/REQUIREMENTS.md を参照。

## 技術スタック
- データ: `data/skills.yaml`(単一ソース。README も将来のサイトもここから生成)
- 生成スクリプト: Go
- CI: GitHub Actions(YAML lint + README 生成差分チェック)
- サイト: なし(M3 で Astro + Cloudflare 予定)

## ディレクトリ構成
- `data/skills.yaml` — 掲載スキルの単一ソース(スキーマはファイル冒頭コメント)
- `cmd/genreadme/` — README 生成スクリプト(Go)
- `README.md` — 生成物(直接編集禁止)
- `docs/` — IDEA.md / REQUIREMENTS.md

## コマンド
- README 生成: `go run ./cmd/genreadme`
- ビルド確認: `go build ./...`

## 規約・注意点
- **README.md は直接編集しない**。必ず `data/skills.yaml` を編集して生成する(手編集は次回生成で消える)
- リスト本文・説明文は英語で書く(掲載スキル自体が日本語でも、description は英語)
- 掲載基準: ネタ系・エンタメ系であること。実務・生産性系スキルは対象外
- 掲載対象は GitHub リポジトリのみ(スキル/プラグイン/Claude Code 連動ツール)。記事(Zenn/Qiita 等)のみのスキルは載せない。記事にリポジトリがあればリポジトリを載せる
