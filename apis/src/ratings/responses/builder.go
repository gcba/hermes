package responses

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	Type        int    `json:"type"`
	Description string `json:"description"`
}
