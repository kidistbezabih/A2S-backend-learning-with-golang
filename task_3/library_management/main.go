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
	library := services.NewLibrary()

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
		 

		case "2":
			var id int64
			fmt.Print("Enter book ID to remove: ")
			fmt.Scanln(&id)
			library.RemoveBook(id)

		case "3":
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

		case "4":
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

		case "5":
			books := library.ListAvailableBooks()
			fmt.Println("Available Books:")
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.Id, book.Title, book.Author, book.Status)
			}
		
		case "6":
			var memberID int64
			fmt.Print("Enter member ID to list borrowed books: ")
			fmt.Scanln(&memberID)
			books := library.ListBorrowedBooks(memberID)
			fmt.Println("Borrowed Books:")
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.Id, book.Title, book.Author, book.Status)
			}

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}