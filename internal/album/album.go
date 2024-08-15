package album

type album struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	Artist artist
}
