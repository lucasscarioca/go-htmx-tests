package models

type Playlist struct {
	Id          string // ObjectId
	Title       string
	Description string
	CoverURL    string
	Categories  []Category
	Tabs        []Tab // []ObjectId
	Public      bool
	Views       uint64
	CreatedAt   string
}
