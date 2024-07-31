package services

import (
	"library_management/models"
	"errors"
)


type LibraryManagement interface{
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}


type Library struct{
	Books map[int64]models.Book
	Member map[int64]models.Member
}

// for creating and initialiing new library

func NewLibrary() *Library{
	return &Library{
		Books: make(map[int64]models.Book),
		Member: make(map[int64]models.Member),
	}
}

// methods

func (l *Library) AddBook(book models.Book){
	l.Books[book.Id] = book
}

func (l *Library) RemoveBook(bookID int64){
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int64, memberID int64) error{
	book, exist := l.Books[bookID]

	if !exist{
		return errors.New("Book doesn't exist")
	}

	if book.Status == "borrowed"{
		return errors.New("Book is not available")
	}

	member, exist := l.Member[memberID]
	if !exist{
		return errors.New("Member does not exist")
	}

	book.Status = "borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)

	return nil
}

func (l *Library) ReturnBook(bookID int64, memberID int64) error{
	book, exist := l.Books[bookID]

	if !exist {
		return errors.New("Book was not exist")
	}

	if book.Status == "available"{
		return errors.New("The Book was not borrowed")
	}

	member, exist := l.Member[memberID]

	if !exist{
		return errors.New("member does not exist")
	}

	book.Status = "available"

	for i, book := range member.BorrowedBooks{
		if book.Id == bookID{
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	return nil
}

func (l *Library)	ListAvailableBooks() []models.Book{
	var availableBooks []models.Book

	for _, book := range l.Books{
		if book.Status == "available"{
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}


func (l *Library) ListBorrowedBooks(memberID int64) []models.Book{
	member, exist := l.Member[memberID]

	if !exist{
		return nil
	}

	var borrowedBooksList []models.Book

	for _, id := range member.BorrowedBooks{
		borrowedBooksList = append(borrowedBooksList, id)
	}
return borrowedBooksList

}