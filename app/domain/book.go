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
	GetBookByCriteria(search SearchCriteria) (result []Book, err error)
	GetAll() (result []Book, err error)
	GetById(bookID string) (result Book, err error)
	Create(book BookCreateRequest, filename string) (result Book, err error)
	Update(bookID string, book BookUpdateRequest, filename string) (result Book, err error)
	Delete(bookID string, filename string) error
}

type BookRepository interface {
	Search(search SearchCriteria) (result []Book, err error)
	GetAll() (result []Book, err error)
	GetById(bookID string) (result Book, err error)
	Put(book Book, filename string) (result Book, err error)
	Delete(bookID string, filename string) error
}
