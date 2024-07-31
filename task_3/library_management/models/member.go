package models

type Member struct{
	Id int64
	Name string
	BorrowedBooks []Book
}
