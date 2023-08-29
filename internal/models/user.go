package models

type User struct {
	Id              string // ObjectId
	Name            string
	Email           string
	Password        string
	Avatar          string
	MidiaLinks      []string
	PinnedPlaylists []Playlist // []ObjectId
	PinnedTabs      []Tab      // []ObjectId
	Groups          []Group    // []ObjectId
	Public          bool
	Views           uint64
	CreatedAt       string
}

func NewUser() User {
	return User{}
}
