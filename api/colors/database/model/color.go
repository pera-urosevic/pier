package model

type Color struct {
	Name string `json:"name"`
	H    int64  `json:"h"`
	S    int64  `json:"s"`
	L    int64  `json:"l"`
}
