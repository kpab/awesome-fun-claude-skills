// genreadme generates README.md from data/skills.yaml.
// data/skills.yaml が唯一のソース。README.md を直接編集しないこと。
package main

import (
	"fmt"
	"os"
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
	// TODO(M1): gopkg.in/yaml.v3 で data/skills.yaml を読み込み、
	// カテゴリ順・スキル名順で README.md を生成する。
	// バリデーション: category 参照切れ / URL 重複 / 必須フィールド欠落はエラーにする。
	fmt.Fprintln(os.Stderr, "genreadme: not implemented yet (see TODO in cmd/genreadme/main.go)")
	os.Exit(1)
}
