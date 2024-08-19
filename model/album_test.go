package model

import "testing"

var noAlbums = []Album{}

func TestGet(t *testing.T) {
	len := len(albums)
	if len == 0 {
		t.Fatal("album doesn't include any elements")
	}
	t.Logf("%v elements found", len)
}

func TestGetFils(t *testing.T) {
	len := len(noAlbums)
	if len == 0 {
		t.Logf("album doesn't include any elements")
	}
	t.Logf("%v elements found", len)
}

func TestPost(t *testing.T) {
	var album = Album{
		Id:    3,
		Title: "Black Album",
		Price: 35,
		Artist: artist{
			Id:   1,
			Name: "Jayz",
		},
	}
	albums = append(albums, album)
	t.Logf("the new albums count is %v", len(albums))
}
