---
name: collect
description: ネタ系スキルの新規候補を発掘する。「新しいネタスキル探して」「候補集めて」「巡回して」と言われたら使う。GitHub/Zenn/Qiita/Hacker News を定型クエリで巡回し、未掲載の候補をリストアップする。
---

# ネタ系スキルの候補を発掘する

1. `data/skills.yaml` を読み、掲載済み URL の一覧を把握する(重複提案を防ぐ)
2. 以下のソースを巡回する(全部でなくてよい。前回から日が浅いソースはスキップ可):
   - **GitHub**: `gh search repos "<kw>" --sort stars --limit 15 --json fullName,description,stargazersCount,updatedAt`
     - キーワード例: `claude persona`, `claude personality`, `claude skill fun`, `claude joke`, `claude character`
     - 新着を狙うなら `--created ">直近2ヶ月の日付"` を付けてもう1周
   - **Zenn**: WebFetch で `https://zenn.dev/topics/claudecode/feed` を確認し、キャラ化・口調・ネタ系の記事を拾う
   - **Qiita / Web全般**: WebSearch で「Claude Code キャラ 口調 スキル」「Claude Code 面白い ネタ」等
   - **Hacker News**: WebFetch で `https://hn.algolia.com/api/v1/search?query=claude%20persona&tags=story`
3. 候補の判定基準:
   - ネタ系・エンタメ系であること(実務・生産性系は除外)
   - 実体があること(SKILL.md やプロンプトが公開されている。紹介だけの記事は、再現可能な設定が載っていれば可)
   - 未掲載であること(手順1の一覧と突き合わせ)
4. 候補を表形式でユーザーに提示する: スキル名 / URL / 種別(リポジトリ・記事) / スター数or反応 / 一言説明 / おすすめ度
5. ユーザーが選んだものを add-skill の手順で追加する(1件ずつ)
