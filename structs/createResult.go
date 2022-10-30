package structs

type CreateResult struct {
	Error   Error   `json:"erro"`
	Article Article `json:"article"`
}
