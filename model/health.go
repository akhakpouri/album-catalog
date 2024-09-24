package model

type Health struct {
	Status  string `json:"status"`
	Details []Detail
}
