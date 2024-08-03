package main

import (
	"fmt"
	"library_management/models"
	"library_management/services"
	"bufio"
	"os"
	"strings"
)

func main() {
	inmemorydb := services.NewLibrary()
	postgresdb := services.NewPostgresLibrary()
	var library services.LibraryManagement

	if os.Getenv("DEPLOYMENT_TYPE") == "production"{
		library = postgresdb
	}else{
		library = inmemorydb
	}

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")

		fmt.Print("Select an option: ")
		reader := bufio.NewReader(os.Stdin)

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			addBook(library)		 

		case "2":
			removeBookFromLib(library)

		case "3":
			borrowFrom(library)

		case "4":
			returnBookTo(library)

		case "5":
			availableBooksIn(library)
		
		case "6":
			listOfBorrowedBooksFrom(library)

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}}

func addBook(library services.LibraryManagement){
	var id int64
			var title, author string
			var status string
			fmt.Print("Enter book ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter book title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter book author: ")
			fmt.Scanln(&author)
			fmt.Print("Enter book status (available/borrowed): ")
			fmt.Scanln(&status)

			book := models.Book{Id:id , Title: title, Author : author, Status : status}
			library.AddBook(book)
}

func removeBookFromLib(library services.LibraryManagement){
			var id int64
			fmt.Print("Enter book ID to remove: ")
			fmt.Scanln(&id)
			library.RemoveBook(id)
}

func borrowFrom(library services.LibraryManagement){
			var bookID, memberID int64
			fmt.Print("Enter book ID to borrow: ")
			fmt.Scanln(&bookID)
			fmt.Print("Enter member ID: ")
			fmt.Scanln(&memberID)
			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed successfully.")
			}

}

func returnBookTo(library services.LibraryManagement){
	var bookID, memberID int64
			fmt.Print("Enter book ID to return: ")
			fmt.Scanln(&bookID)
			fmt.Print("Enter member ID: ")
			fmt.Scanln(&memberID)
			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned successfully.")
			}
}


func availableBooks(library services.LibraryManagement){
	books := library.ListAvailableBooks()
			fmt.Println("Available Books:")
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.Id, book.Title, book.Author, book.Status)
			}
}


func listOfBorrowedBooksFrom(library services.LibraryManagement){
	var memberID int64
			fmt.Print("Enter member ID to list borrowed books: ")
			fmt.Scanln(&memberID)
			books := library.ListBorrowedBooks(memberID)
			fmt.Println("Borrowed Books:")
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.Id, book.Title, book.Author, book.Status)
			}
}