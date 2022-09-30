package entity

type Food struct {
	Id    int     `json:"id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
	Type  int     `json:"type"`
}
