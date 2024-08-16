package album

import (
	"errors"
	"fmt"
	"slices"
)

type Album struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	Artist artist  `json:"artist"`
}

var albums = []Album{
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

func GetAlbums() []Album {
	return albums
}

func GetAlbumById(id int) (Album, error) {
	idx := slices.IndexFunc(albums, func(a Album) bool {
		return a.Id == id
	})
	if idx < 0 {
		err := fmt.Sprintf("Can't find an album with id %v", id)
		return Album{}, errors.New(err)
	}

	return albums[idx], nil
}

func AddAlbum(a Album) {
	albums = append(albums, a)
}
