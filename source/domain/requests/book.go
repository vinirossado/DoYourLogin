package requests

type BookRequest struct {
	AuthorId    uint    `json:"author_id" binding:"required" example:"1"`
	Name        string  `json:"name" binding:"required" example:"NameOfYourBook"`
	ISBN        string  `json:"isbn" binding:"required" example:"ISBNOfYourBook"`
	Description string  `json:"description" example:"DescriptionOfYourBook"`
	Price       float64 `json:"price" example:"1970"`
	ReleaseYear int     `json:"release_year" example:"1970"`
	Categories  []uint  `json:"categories" binding:"required" example:"[1,2,3]"`
}
