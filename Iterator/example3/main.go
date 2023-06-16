package main

import (
	"fmt"
	"sort"
)

type Book struct {
	Title string
}

func NewBook(title string) *Book {
	return &Book{title}
}

func (b *Book) GetTitle() string {
	return b.Title
}

type Books []*Book

func (b Books) Len() int {
	return len(b)
}

func (b Books) Less(i, j int) bool {
	return b[i].Title < b[j].Title
}

func (b Books) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

type BookShelf struct {
	Books
}

func NewBookShelf() *BookShelf {
	return &BookShelf{}
}

func (bs *BookShelf) AddBook(book *Book) {
	bs.Books = append(bs.Books, book)
}

func (bs *BookShelf) Iterator() Iterator {
	return &BookIterator{Books: bs.Books, Index: 0}
}

func (bs *BookShelf) SortedIterator() Iterator {
	SortedBooks := make(Books, len(bs.Books))
	copy(SortedBooks, bs.Books)
	sort.Sort(SortedBooks)
	return &SortedBookIterator{BookIterator{Books: SortedBooks, Index: 0}}
}

type Iterator interface {
	HasNext() bool
	Next() *Book
}

type BookIterator struct {
	Books []*Book
	Index int
}

func (bi *BookIterator) HasNext() bool {
	return bi.Index < len(bi.Books)
}

func (bi *BookIterator) Next() *Book {
	if bi.HasNext() {
		book := bi.Books[bi.Index]
		bi.Index++
		return book
	}
	return nil
}

type SortedBookIterator struct {
	BookIterator
}

func main() {
	bookshelf := NewBookShelf()
	bookshelf.AddBook(NewBook("Book 3"))
	bookshelf.AddBook(NewBook("Book 2"))
	bookshelf.AddBook(NewBook("Book 1"))

	iterator := bookshelf.Iterator()
	for iterator.HasNext() {
		book := iterator.Next()
		fmt.Println(book.GetTitle())
	}

	iterator = bookshelf.SortedIterator()
	for iterator.HasNext() {
		book := iterator.Next()
		fmt.Println(book.GetTitle())
	}
}
