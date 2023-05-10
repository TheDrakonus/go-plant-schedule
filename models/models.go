package models

type Plant struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Plants []Plant

type SampleData struct {
	Plants []Plant `json:"plants"`
}
