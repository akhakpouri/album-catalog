package album

import (
	"errors"
	"fmt"
	"slices"
)

type album struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	Artist artist  `json:"artist"`
}

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

func GetAlbums() []album {
	return albums
}

func GetAlbumById(id int) (album, error) {
	idx := slices.IndexFunc(albums, func(a album) bool {
		return a.Id == id
	})
	if idx < 0 {
		err := fmt.Sprintf("Can't find an album with id %v", id)
		return album{}, errors.New(err)
	}

	return albums[idx], nil
}

func AddAlbum(a album) {
	albums = append(albums, a)
}
