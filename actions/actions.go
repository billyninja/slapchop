package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/billyninja/slapchop/chopper"
	"github.com/billyninja/slapchop/puzzler"
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	// 3rd party
	"github.com/hoisie/mustache"
	"github.com/julienschmidt/httprouter"
)

/* Constants */

type ActionsConfig struct {
	HostName      string
	Port          string
	Host          string
	UploadDir     string
	MaxUploadSize int64
	TileSize      int
	PuzzlerHost   string
	TemplatePath  string
}

/* Let's use here the CRUD standard names */
func (ac *ActionsConfig) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	username := ps.ByName("username")
	r.ParseMultipartForm(ac.MaxUploadSize)

	// Opening the file
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		log.Println(err)
		return
	}

	// Setting to close filehandler at the end of this function
	defer file.Close()
	chop_id := time.Now().Format("020106150405")
	path := fmt.Sprintf("%s/%s/%s", ac.UploadDir, username, chop_id)
	img, format, err := chopper.Load(file)
	tiles := chopper.Slice(*img, ac.TileSize, format, path)
	chopper.SaveAll(tiles)

	href_base := fmt.Sprintf("http://%s:8080/upload/%s/%s", ac.HostName, username, chop_id)

	tilesR := make([]*chopper.TileEntry, len(tiles))
	for i, t := range tiles {
		tilesR[i] = t.ToResp(href_base)
	}

	resp := chopper.CreateResponse{
		User:   username,
		ChopId: chop_id,
		Href:   fmt.Sprintf("http://%s/chopit/%s/%s", ac.Host, username, chop_id),
		Tiles:  tilesR,
	}

	// If it is setted to connect with Puzzler Service
	if len(ac.PuzzlerHost) > 0 {
		status, body, err := puzzler.CreatePuzzle(ac.PuzzlerHost, username, tilesR)
		if err != nil {
			log.Print(err)
		}

		if status == 201 {
			puzzler_resp := puzzler.CreateResponse{}
			_ = json.Unmarshal(body, &puzzler_resp)
			resp.PuzzleHref = puzzler_resp.PuzzleHref
			resp.SolutionHref = puzzler_resp.SolutionHref
		}
	}

	json_resp, err := json.Marshal(&resp)
	if err != nil {
		log.Fatalf("errr %v", err)
	}

	w.Write(json_resp)
	w.WriteHeader(http.StatusOK)
}

func (ac *ActionsConfig) ReadAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	username := ps.ByName("username")
	path := fmt.Sprintf("%s/%s", ac.UploadDir, username)
	dirs, _ := ioutil.ReadDir(path)
	slapchops := make([]*chopper.SlapchopEntry, len(dirs))
	for i, d := range dirs {
		slapchops[i] = chopper.NewSlapchop(ac.Host, username, d.Name())
	}
	resp := chopper.ReadAllResponse{
		User:      username,
		Slapchops: slapchops,
	}

	w.WriteHeader(http.StatusOK)
	json_resp, _ := json.Marshal(&resp)
	w.Write(json_resp)
}

func (ac *ActionsConfig) Read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	username := ps.ByName("username")
	chopid := ps.ByName("chopid")
	s := chopper.NewSlapchop(ac.Host, username, chopid)
	t_files, _ := s.LoadFiles(ac.UploadDir)
	tiles := s.LoadTiles(ac.HostName, t_files)

	resp := chopper.ReadResponse{
		User:  username,
		Id:    chopid,
		Tiles: tiles,
	}

	w.WriteHeader(http.StatusOK)
	json_resp, _ := json.Marshal(&resp)
	w.Write(json_resp)
}

func (ac *ActionsConfig) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := ps.ByName("username")
	chopid := ps.ByName("chopid")

	path := fmt.Sprintf("%s/%s/%s", ac.UploadDir, username, chopid)
	err := os.RemoveAll(path)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s %s successfully deleted", username, chopid)))
}

func (ac *ActionsConfig) Preview(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	username := ps.ByName("username")
	chopid := ps.ByName("chopid")
	s := chopper.NewSlapchop(ac.Host, username, chopid)
	t_files, _ := s.LoadFiles(ac.UploadDir)
	tiles := s.LoadTiles(ac.HostName, t_files)
	grid := s.Grid(tiles)

	html := mustache.RenderFile(ac.TemplatePath, grid)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func (ac *ActionsConfig) Random(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	username := ps.ByName("username")
	chopid := ps.ByName("chopid")
	s := chopper.NewSlapchop(ac.Host, username, chopid)
	t_files, _ := s.LoadFiles(ac.UploadDir)
	tiles := s.LoadTiles(ac.HostName, t_files)
	grid := s.Grid(tiles)
	grid = s.ShuffleGrid(grid)

	html := mustache.RenderFile(ac.TemplatePath, grid)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func (ac *ActionsConfig) Pixelette(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	username := ps.ByName("username")
	chopid := ps.ByName("chopid")
	s := chopper.NewSlapchop(ac.Host, username, chopid)
	t_files, _ := s.LoadFiles(ac.UploadDir)

	up := s.UploadPoint(ac.UploadDir)
	var sR, sG, sB, sA uint64
	sR, sG, sB, sA = 0, 0, 0, 0
	for _, f := range t_files {
		file, _ := os.Open(fmt.Sprintf("%s/%s", up, f.Name()))
		img, _, _ := chopper.Load(file)
		subImg := image.NewRGBA(image.Rect(0, 0, ac.TileSize, ac.TileSize))
		draw.Draw(subImg, subImg.Bounds(), *img, image.Rect(0, 0, ac.TileSize, ac.TileSize).Min, draw.Src)

		for x := 0; x <= ac.TileSize; x++ {
			for y := 0; y <= ac.TileSize; y++ {
				r, g, b, a := subImg.At(x, y).RGBA()
				sR += uint64(r)
				sG += uint64(g)
				sB += uint64(b)
				sA += uint64(a)
			}
		}
		ms := uint64(ac.TileSize * ac.TileSize)
		println(sR/ms, sG/ms, sB/ms, sA/ms)

	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(``))
}

//TEST HELPER: Creates a new file upload http request with optional extra params
func NewfileUploadRequest(uri string, paramName, path string) (*http.Request, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	return request, err
}
