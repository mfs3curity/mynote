package dto

type CreatePost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	SectionName string `json:"section_name"`
	UserId      int
}
type ResponsePost struct {
	Username    string `json:"username"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SectionName string `json:"section_name"`
}
