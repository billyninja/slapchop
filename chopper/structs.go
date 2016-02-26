package chopper

import (
	"fmt"
	"io/ioutil"
	"os"
)

type SlapchopEntry struct {
	Id   string `json:"id"`
	User string `json:"user"`
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
	User         string       `json:"user"`
	ChopId       string       `json:"chopid"`
	Href         string       `json:"href"`
	PuzzleHref   string       `json:"puzzle_href"`
	SolutionHref string       `json:"solution_href"`
	Tiles        []*TileEntry `json:"tiles"`
}

func NewSlapchop(host_str, username, id string) *SlapchopEntry {

	href := fmt.Sprintf("http://%s/chopit/%s/%s", host_str, username, id)

	return &SlapchopEntry{
		Id:   id,
		User: username,
		Href: href,
	}
}

func (s *SlapchopEntry) LoadFiles(UploadDir string) ([]os.FileInfo, error) {
	path := fmt.Sprintf("%s/%s/%s", UploadDir, s.User, s.Id)
	files, err := ioutil.ReadDir(path)
	return files, err
}

func (s *SlapchopEntry) LoadTiles(host string, files []os.FileInfo) []*TileEntry {
	tiles := make([]*TileEntry, len(files))

	for i, f := range files {
		fname := f.Name()
		href := fmt.Sprintf("http://%s:8080/upload/%s/%s/%s", host, s.User, s.Id, fname)
		tiles[i] = &TileEntry{
			Filename: fname,
			Href:     href,
		}
	}

	return tiles
}
