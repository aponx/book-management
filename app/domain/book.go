package domain

type SearchCriteria struct {
	Title     *string `col:"title"`
	Author    *string `col:"author"`
	Publisher *string `col:"publisher"`
}

type Book struct {
	BookID    string `json:"bookID"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Year      string `json:"year"`
	Qty       int    `json:"qty"`
	Out       int    `json:"out"`
}

type BookCreateRequest struct {
	BookID    string `json:"bookID"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Year      string `json:"year"`
	Qty       int    `json:"qty"`
}

type BookUpdateRequest struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Year      string `json:"year"`
	Qty       int    `json:"qty"`
	Out       int    `json:"out"`
}

type BookUsecase interface {
	GetAll() (result []Book, err error)
	GetById(bookID string) (result Book, err error)
	Create(book *BookCreateRequest) (result *Book, err error)
	Update(bookID string, book *BookUpdateRequest) (result *Book, err error)
	Delete(bookID string) (err error)
}
