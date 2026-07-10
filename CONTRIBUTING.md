# Contributing

Thanks for helping make this list more fun! Submissions are welcome — as long as they make someone smile.

## Ground rules

- **Fun-first only.** Personas, jokes, games, desktop pets, vibes. Productivity and dev-workflow skills belong in the other awesome lists — a game *development* toolkit is work; something you *play with* is fun.
- **GitHub repos only.** Skills, plugins, or Claude Code-integrated toys published as a GitHub repository. Article-only skills (Zenn/Qiita/blog posts) are not listed; if the article has a repo, link the repo.
- **Quality first.** The repo must have real content (a SKILL.md, plugin structure, prompts, or working code). Among similar entries, higher-starred and better-maintained ones win.

## How to add a skill

1. Edit [`data/skills.yaml`](data/skills.yaml). **Never edit `README.md` directly** — it is generated and your change will be overwritten.
2. Add your entry under `skills:` in the right category, keeping the in-category order (curation order: roughly stars-descending).
3. Regenerate the README:

   ```sh
   go run ./cmd/genreadme
   ```

   The generator also validates the YAML (required fields, category references, duplicate URLs) and will tell you if something is wrong.
4. Commit **both** `data/skills.yaml` and `README.md`, then open a pull request.

## Entry schema

```yaml
- name: display name
  url: https://github.com/owner/repo   # GitHub repository URL
  source: repo                          # always "repo" for new entries
  category: personas                    # one of categories[].id
  language: en                          # primary language of the repo/README (en, ja, zh, ko, ...)
  description: 1-2 sentences in English; make the "fun" obvious.
  author: owner                         # GitHub handle (optional)
  tags: [optional, keywords]            # optional
```

Descriptions are always in English, even when the skill itself is in another language — the `language` field adds a *(in Japanese)*-style note automatically. Star counts are rendered as live badges, so don't put them in the YAML.

## Adding a category

Categories live at the top of `data/skills.yaml`; their order there is the section order in the README. Open an issue first if you think a new category is needed — we prefer few, well-filled sections.

## CI

Every pull request runs a check that `data/skills.yaml` validates and `README.md` matches what the generator produces. If CI fails with "README.md is out of date", run `go run ./cmd/genreadme` and commit the result.
