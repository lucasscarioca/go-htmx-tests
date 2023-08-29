package models

type Tab struct {
	Id         string // ObjectId
	Title      string
	Author     string // ObjectId
	Categories []Category
	MidiaLinks []string
	CoverURL   string
	Tuning     string
	Content    string
	Public     bool
	Views      uint64
	CreatedAt  string
}
