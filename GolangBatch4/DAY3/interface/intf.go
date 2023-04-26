package main

import "fmt"

type Test interface {
	Print()
}

type Book struct {
	BookId int
	Title  string
	Author string
}

func (b Book) Print() {
	fmt.Println("Print from Book")
	fmt.Println(b)

}

func (b Book) String() string {
	fmt.Println("This is String from Book type")
	return fmt.Sprintln("BookID:", b.BookId, "Title:", b.Title, "Author:", b.Author)
}

type Store struct {
	StoreID int
	Name    string
	Address string
}

func (s Store) Print() {
	fmt.Println("Print from store")
	fmt.Println(s)
}

func (s Store) String() string {
	fmt.Println("This is String from Store type")
	return fmt.Sprintln("StoreID:", s.StoreID, "Name:", s.Name, "Address:", s.Address)
}

func main() {

	var obj1 Test

	store := Store{StoreID: 1, Name: "My Store", Address: "Mfar"}

	book := Book{BookId: 2, Title: "Golang", Author: "Google"}

	obj1 = store

	obj1.Print()

	obj1 = book

	obj1.Print()

	fmt.Println(store)
	fmt.Println(book)

}
