package dto

import "time"

type CreatSection struct {
	Name string `json:"name" binding:"required"`
}
type ResponseSection struct {
	Name             string      `json:"name"`
	TitlePost        []string    `json:"title_post"`
	CreatedAtSection time.Time   `json:"created_section"`
	CreatedAtPost    []time.Time `json:"created_post"`
	IdPost           []int       `json:"id_post"`
}
