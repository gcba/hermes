package responses

type Post struct {
	Meta Meta `json:"meta"`
}

func PostResponse(status int) *Post {

}
