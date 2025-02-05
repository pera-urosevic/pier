package types

type TVMazeSearchResult struct {
	Title   string  `json:"title"`
	Label   string  `json:"label"`
	ID      int64   `json:"id"`
	Text    string  `json:"text"`
	URL     string  `json:"url"`
	Image   *string `json:"image"`
	Year    string  `json:"year"`
	Status  string  `json:"status"`
	Runtime *int64  `json:"runtime"`
}
