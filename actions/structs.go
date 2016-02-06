package actions

type SlapchopEntry struct {
	Id string
	Href string
}

type TileEntry struct {
	Filename string `json:"filename"`
	Href string `json:"href"`
}

type ReadResponse struct {
	User string `json:"user"`
	Id string `json:"id"`
	Tiles []*TileEntry `json:"tiles"`
}

type ReadAllResponse struct {
	User string
	Slapchops []*SlapchopEntry
}
