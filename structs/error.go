package structs

type Error struct {
	Status  int16  `json:"status"`
	Message string `json:"message"`
}
