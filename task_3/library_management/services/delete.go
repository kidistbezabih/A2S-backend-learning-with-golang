package services


type PostgresLibrary struct{
	instance *postgresclient
}

// for creating and initialiing new library

func NewPostgresLibrary() *PostgresLibrary{
	postgresLink := os.Getenv("POSTGRES_LINK")
	// some logic to connect to postgres
	postgresclient := newasdflasdkf
	return &PostgresLibrary{
		postgresclient: postgresclient
	}
}

// methods

func (l *Library) AddBook(book models.Book){
	// some code specific to postgres to add a book
}

func (l *Library) RemoveBook(bookID int64){
	// some code specific to postgres to delete a book
}

func (l *Library) BorrowBook(bookID int64, memberID int64) error{
	// some code specific to postgres to borrow a book

	return nil
}

func (l *Library) ReturnBook(bookID int64, memberID int64) error{
	// some code specific to postgres to return a book

	return nil
}

func (l *Library)	ListAvailableBooks() []models.Book{
	var availableBooks []models.Book
	// some code specific to postgres to add list of books to availableBooks
	return availableBooks
}


func (l *Library) ListBorrowedBooks(memberID int64) []models.Book{
	var borrowedBooksList []models.Book

	// some code specific to postgres to add list of books to availableBooks
return borrowedBooksList

}