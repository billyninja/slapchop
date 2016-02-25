package chopper

type SlapchopEntry struct {
	Id   string `json:"id"`
	Href string `json:"href"`
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
	User      string           `json:"user"`
	Slapchops []*SlapchopEntry `json:"slapchops"`
}

type CreateResponse struct {
	User   string       `json:"user"`
	ChopId string       `json:"chopid"`
	Href   string       `json:"href"`
	Tiles  []*TileEntry `json:"tiles"`
}
