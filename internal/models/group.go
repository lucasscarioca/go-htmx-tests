package models

type Group struct {
	Id              string // ObjectId
	Name            string
	Description     string
	Avatar          string
	MidiaLinks      []string
	PinnedPlaylists []Playlist // []ObjectId
	PinnedTabs      []Tab      // []ObjectId
	Users           []User     // []ObjectId
	Public          bool
	Views           uint64
	CreatedAt       string
}
