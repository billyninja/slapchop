package chopper

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

func (s *SlapchopEntry) Grid(tiles []*TileEntry) [][40]string {
	/* Organize the tiles in a matrix respecting the original organization
	 */
	g := make([][40]string, 40)
	for _, t := range tiles {
		cordStr := strings.Split(strings.Split(t.Filename, ".")[0], "_")
		pX, _ := strconv.Atoi(cordStr[0])
		pY, _ := strconv.Atoi(cordStr[1])
		g[pX][pY] = t.Href
	}

	return g
}

func (s *SlapchopEntry) ShuffleGrid(grid [][40]string) [][40]string {
	boundY := 1
	boundX := 1
	i := 20

	for i > 0 {

		for x, _ := range grid {
			for y, _ := range grid[x] {
				if grid[x][y] == "" {
					break
				}

				if y > boundY {
					boundY = y
				}
				if x > boundX {
					boundY = x
				}
				rx := rand.Intn(boundX)
				ry := rand.Intn(boundY)
				temp := grid[x][y]
				grid[x][y] = grid[rx][ry]
				grid[rx][ry] = temp
			}
		}
		i--
	}

	return grid
}
