package chopper

type SlapchopEntry struct {
	Id   string
	Href string
}

type TileEntry struct {
	Filename string `json:"-"`
	Href     string `json:"href"`

	PosX int `json:"pos_x"`
	PosY int `json:"pos_y"`
	AbsX int `json:"abs_x"`
	AbsY int `json:"abs_y"`
}

type ReadResponse struct {
	User  string       `json:"user"`
	Id    string       `json:"id"`
	Tiles []*TileEntry `json:"tiles"`
}

type DeleteResponse struct {
	User  string       `json:"user"`
	Id    string       `json:"id"`
	Tiles []*TileEntry `json:"tiles"`
}

type ReadAllResponse struct {
	User      string
	Slapchops []*SlapchopEntry
}

type CreateResponse struct {
	User   string
	ChopId string
	Href   string
	Tiles  []*TileEntry
}
