package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Book represents a book in the library
type Book struct {
	Title        string
	Author       string
	ISBN         string
	IsCheckedOut bool
}

// Library represents a collection of books
type Library struct {
	Books []Book
}

// LibraryActions defines actions that can be performed on a library
type LibraryActions interface {
	AddBook(title, author, isbn string)
	CheckOutBook(isbn string) error
	ReturnBook(isbn string) error
	ListAvailableBooks()
}

// AddBook adds a new book to the library
func (l *Library) AddBook(title, author, isbn string) {
	l.Books = append(l.Books, Book{Title: title, Author: author, ISBN: isbn, IsCheckedOut: false})
	fmt.Printf("Added \"%s\" by %s to the library\n", title, author)
}

// CheckOutBook checks out a book by its ISBN
func (l *Library) CheckOutBook(isbn string) error {
	for i, book := range l.Books {
		if book.ISBN == isbn {
			if book.IsCheckedOut {
				return errors.New("book is already checked out")
			}
			l.Books[i].IsCheckedOut = true
			fmt.Printf("You checked out \"%s\"\n", book.Title)
			return nil
		}
	}
	return errors.New("book not found")
}

// ReturnBook returns a book by its ISBN
func (l *Library) ReturnBook(isbn string) error {
	for i, book := range l.Books {
		if book.ISBN == isbn {
			if !book.IsCheckedOut {
				return errors.New("book was not checked out")
			}
			l.Books[i].IsCheckedOut = false
			fmt.Printf("You returned \"%s\"\n", book.Title)
			return nil
		}
	}
	return errors.New("book not found")
}

// ListAvailableBooks lists all available (not checked out) books in the library
func (l *Library) ListAvailableBooks() {
	fmt.Println("Available Books:")
	for _, book := range l.Books {
		if !book.IsCheckedOut {
			fmt.Printf("- %s by %s\n", book.Title, book.Author)
		}
	}
}

// PromptBookDetails asks the user for book details and adds it to the library
func PromptBookDetails(l *Library) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter book author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	fmt.Print("Enter book ISBN: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	l.AddBook(title, author, isbn)
}

func main() {
	// Create a library instance
	library := Library{}

	// Prompt user to add books to the library
	for {
		var addMore string
		fmt.Print("Would you like to add a book to the library? (yes/no): ")
		fmt.Scanln(&addMore)
		if strings.ToLower(addMore) != "yes" {
			break
		}
		PromptBookDetails(&library)
	}

	// List all available books
	library.ListAvailableBooks()

	// Check out a book by ISBN
	var isbn string
	fmt.Print("Enter ISBN to check out a book: ")
	fmt.Scanln(&isbn)
	err := library.CheckOutBook(isbn)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Return a book by ISBN
	fmt.Print("Enter ISBN to return a book: ")
	fmt.Scanln(&isbn)
	err = library.ReturnBook(isbn)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// List available books again
	library.ListAvailableBooks()
}
