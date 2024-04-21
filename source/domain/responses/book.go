package responses

import "time"

type BookResponse struct {
	AuthorId    uint      `json:"author_id" binding:"required" example:"1"`
	Name        string    `json:"name" example:"NameOfYourBook"`
	ISBN        string    `json:"isbn" example:"ISBNOfYourBook"`
	Description string    `json:"description" example:"DescriptionOfYourBook"`
	Price       float64   `json:"price" example:"1970"`
	ReleaseYear time.Time `json:"release_year" example:"1970"`

	Categories []CategoryResponse `json:"categories"`
}
