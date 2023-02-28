package model

type Notecard struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Entry string `json:"entry"`
}

type Section struct {
	ID       uint64      `json:"id"`
	Title    string      `json:"title"`
	Subtitle string      `json:"subtitle"`
	Notes    []*Notecard `json:"notecard"`
}
