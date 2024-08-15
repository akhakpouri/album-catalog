package album

import "testing"

var albums = []album{
	{
		Id:    1,
		Title: "Reasonable Doubt",
		Price: 20.00,
		Artist: artist{
			Id:   1,
			Name: "JayZ",
		},
	},
	{
		Id:    2,
		Title: "Port Of Miami",
		Price: 10.00,
		Artist: artist{
			Id:   2,
			Name: "Rick Ross",
		},
	},
}

var noAlbums = []album{}

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
	var album = album{
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
