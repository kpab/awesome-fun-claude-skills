// genreadme generates README.md from data/skills.yaml.
// data/skills.yaml が唯一のソース。README.md を直接編集しないこと。
package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	dataPath   = "data/skills.yaml"
	readmePath = "README.md"
)

// Category は README のセクションに対応する。
type Category struct {
	ID          string `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

// Skill は掲載スキル1件。スキーマの詳細は data/skills.yaml 冒頭のコメントを参照。
type Skill struct {
	Name        string   `yaml:"name"`
	URL         string   `yaml:"url"`
	Source      string   `yaml:"source"`
	Category    string   `yaml:"category"`
	Language    string   `yaml:"language"`
	Description string   `yaml:"description"`
	Author      string   `yaml:"author"`
	Tags        []string `yaml:"tags"`
}

// Data は data/skills.yaml 全体。
type Data struct {
	Categories []Category `yaml:"categories"`
	Skills     []Skill    `yaml:"skills"`
}

func main() {
	raw, err := os.ReadFile(dataPath)
	if err != nil {
		fatal(err.Error())
	}
	var data Data
	dec := yaml.NewDecoder(strings.NewReader(string(raw)))
	dec.KnownFields(true)
	if err := dec.Decode(&data); err != nil {
		fatal(fmt.Sprintf("%s: %v", dataPath, err))
	}
	if errs := validate(&data); len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintf(os.Stderr, "genreadme: %s: %s\n", dataPath, e)
		}
		os.Exit(1)
	}
	md := render(&data)
	if err := os.WriteFile(readmePath, []byte(md), 0o644); err != nil {
		fatal(err.Error())
	}
	fmt.Printf("genreadme: wrote %s (%d skills, %d categories)\n", readmePath, len(data.Skills), len(data.Categories))
}

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, "genreadme: "+msg)
	os.Exit(1)
}

func validate(d *Data) []string {
	var errs []string
	catIDs := map[string]bool{}
	for i, c := range d.Categories {
		if c.ID == "" || c.Name == "" {
			errs = append(errs, fmt.Sprintf("categories[%d]: id and name are required", i))
		}
		if catIDs[c.ID] {
			errs = append(errs, fmt.Sprintf("categories[%d]: duplicate id %q", i, c.ID))
		}
		catIDs[c.ID] = true
	}
	urls := map[string]string{}
	for _, s := range d.Skills {
		label := s.Name
		if label == "" {
			label = s.URL
		}
		for field, v := range map[string]string{
			"name": s.Name, "url": s.URL, "source": s.Source,
			"category": s.Category, "language": s.Language, "description": s.Description,
		} {
			if v == "" {
				errs = append(errs, fmt.Sprintf("skill %q: %s is required", label, field))
			}
		}
		if s.Source != "" && s.Source != "repo" && s.Source != "article" {
			errs = append(errs, fmt.Sprintf("skill %q: source must be repo or article, got %q", label, s.Source))
		}
		if s.Category != "" && !catIDs[s.Category] {
			errs = append(errs, fmt.Sprintf("skill %q: unknown category %q", label, s.Category))
		}
		if prev, dup := urls[s.URL]; dup {
			errs = append(errs, fmt.Sprintf("skill %q: duplicate url (already used by %q)", label, prev))
		}
		urls[s.URL] = label
	}
	return errs
}

func render(d *Data) string {
	var b strings.Builder
	b.WriteString("<!-- THIS FILE IS GENERATED from data/skills.yaml — do not edit by hand. -->\n\n")
	b.WriteString("# Awesome Fun Claude Skills\n\n")
	b.WriteString("> A curated list of fun, entertaining, and joke Claude Code skills — personas, humor, and play.\n\n")
	b.WriteString("Everyone curates productivity skills. This list curates the fun ones: skills and plugins that turn Claude Code into a tabletop game master, a desk pet, or your most merciless roast comedian.\n\n")

	b.WriteString("## Contents\n\n")
	for _, c := range d.Categories {
		fmt.Fprintf(&b, "- [%s](#%s)\n", c.Name, anchor(c.Name))
	}
	b.WriteString("\n")

	byCat := map[string][]Skill{}
	for _, s := range d.Skills {
		byCat[s.Category] = append(byCat[s.Category], s)
	}
	// カテゴリ内は data/skills.yaml の記載順のまま出す(キュレーション順)。
	for _, c := range d.Categories {
		fmt.Fprintf(&b, "## %s\n\n", c.Name)
		if c.Description != "" {
			b.WriteString(c.Description + "\n\n")
		}
		for _, s := range byCat[c.ID] {
			fmt.Fprintf(&b, "- [%s](%s)%s — %s%s\n", s.Name, s.URL, starBadge(s), strings.TrimRight(s.Description, "."), suffix(s))
		}
		b.WriteString("\n")
	}

	b.WriteString("## Contributing\n\n")
	b.WriteString("Edit [`data/skills.yaml`](data/skills.yaml) (not this file) and regenerate:\n\n")
	b.WriteString("```sh\ngo run ./cmd/genreadme\n```\n\n")
	b.WriteString("Submissions must be fun-first — personas, jokes, games, vibes. Productivity skills belong in the other awesome lists.\n\n")
	b.WriteString("## License\n\n")
	b.WriteString("List content: [CC0 1.0](LICENSE). Generator code: [MIT](LICENSE-CODE).\n")
	return b.String()
}

// starBadge は github.com のリポジトリエントリに付ける shields.io のスター数バッジを返す。
// スター数はバッジ側で動的に取得されるため YAML に持たない。
func starBadge(s Skill) string {
	if s.Source != "repo" {
		return ""
	}
	path, ok := strings.CutPrefix(s.URL, "https://github.com/")
	if !ok {
		return ""
	}
	path = strings.Trim(path, "/")
	if strings.Count(path, "/") != 1 {
		return ""
	}
	return fmt.Sprintf(" ![stars](https://img.shields.io/github/stars/%s?style=flat-square)", path)
}

// suffix はエントリ末尾の補足(記事か・何語か)を返す。英語のリポジトリなら空。
func suffix(s Skill) string {
	var parts []string
	if s.Source == "article" {
		parts = append(parts, "article")
	}
	if name := langName(s.Language); name != "" {
		parts = append(parts, "in "+name)
	}
	if len(parts) == 0 {
		return "."
	}
	return fmt.Sprintf(". *(%s)*", strings.Join(parts, ", "))
}

func langName(code string) string {
	switch code {
	case "en", "":
		return ""
	case "ja":
		return "Japanese"
	case "zh":
		return "Chinese"
	case "ko":
		return "Korean"
	default:
		return code
	}
}

// anchor は GitHub の見出しアンカー形式(小文字・空白→ハイフン・記号除去)に変換する。
func anchor(heading string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(heading) {
		switch {
		case r >= 'a' && r <= 'z' || r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == ' ' || r == '-':
			b.WriteByte('-')
		}
	}
	return b.String()
}
