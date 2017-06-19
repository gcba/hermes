package responses

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
