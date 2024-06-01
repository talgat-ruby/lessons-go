package library

import (
	"fmt"
)

type Borrowerable interface {
	getBorrowedBy() *Nameable
	setBorrowedBy(*Nameable)
}

type Nameable interface {
	getName() string
}

type Readable interface {
	Nameable
	Borrowerable
}

type Library struct {
	books []Readable
}

type Book struct {
	name       string
	borrowedBy *Nameable
}

type Comics struct {
	collection string
	borrowedBy *Nameable
}

type Student struct {
	name string
}

func (b *Book) getName() string {
	return b.name
}

func (b *Book) getBorrowedBy() *Nameable {
	return b.borrowedBy
}

func (b *Book) setBorrowedBy(n *Nameable) {
	b.borrowedBy = n
}

func (b *Comics) getName() string {
	return b.collection
}

func (b *Comics) getBorrowedBy() *Nameable {
	return b.borrowedBy
}

func (b *Comics) setBorrowedBy(n *Nameable) {
	b.borrowedBy = n
}

func (s *Student) getName() string {
	return s.name
}

func (l *Library) borrowBook(name string, by *Nameable) {
	var borrowedBook Readable

	for _, b := range l.books {
		if b.getName() == name && b.getBorrowedBy() == nil {
			borrowedBook = b
		}
	}

	if borrowedBook == nil {
		fmt.Println("this book already was borrowed")
		return
	}

	borrowedBook.setBorrowedBy(by)
	fmt.Printf("here you book: %s\n", borrowedBook.getName())
}

func (l *Library) returnBook(name string) {
	var borrowedBook Readable

	for _, b := range l.books {
		if b.getName() == name && b.getBorrowedBy() != nil {
			borrowedBook = b
		}
	}

	if borrowedBook == nil {
		fmt.Println("this book is not in our lib or was not borrowed")
		return
	}

	borrowedBook.setBorrowedBy(nil)
	fmt.Printf("Thank you for return: %s", borrowedBook.getName())
}
