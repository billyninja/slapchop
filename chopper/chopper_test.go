package chopper

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	file, _ := os.Open("../test/gopher.jpg")
	img, format, err := Load(file)
	if err != nil {
		t.Error("Something is wrong!")
	}
	println(img, format)
}

func TestSlice(t *testing.T) {
	file, _ := os.Open("../test/gopher.jpg")
	img, format, err := Load(file)
	if err != nil {
		t.Error("Something is wrong!")
	}
	tiles := Slice(*img, 64, format, "/tmp/slapchop/tests")
	if len(tiles) == 0 {
		t.Error("Something is wrong!")
	}
}

func TestSaveAll(t *testing.T) {
	file, _ := os.Open("../test/gopher.jpg")
	img, format, err := Load(file)
	if err != nil {
		t.Error("Something is wrong!")
	}
	tiles := Slice(*img, 64, format, "/tmp/slapchop/tests")
	if len(tiles) == 0 {
		t.Error("Something is wrong!")
	}
	SaveAll(tiles)
}

func TestSaveJpeg(t *testing.T) {
	file, _ := os.Open("../test/gopher.jpg")
	img, format, err := Load(file)
	if err != nil {
		t.Error("Something is wrong!")
	}
	tiles := Slice(*img, 64, format, "/tmp/slapchop/tests")
	if len(tiles) == 0 {
		t.Error("Something is wrong!")
	}
	tiles[0].Save()
}

func TestSavePng(t *testing.T) {
	file, _ := os.Open("../test/gopher.png")
	img, format, err := Load(file)
	if err != nil {
		t.Error("Something is wrong!")
	}
	tiles := Slice(*img, 64, format, "/tmp/slapchop/tests")
	if len(tiles) == 0 {
		t.Error("Something is wrong!")
	}
	tiles[0].Save()
}
