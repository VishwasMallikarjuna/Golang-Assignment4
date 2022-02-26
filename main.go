package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Number  int
	Content string
}

type Book struct {
	Author string
	Pages  []Page
	Name   string
}

type Library struct {
	books []Book
}

func main() {

	Page1Book1 := Page{Number: 1, Content: "chapter 1"}
	Page2Book1 := Page{Number: 2, Content: "chapter 2"}
	Page1Book2 := Page{Number: 1, Content: "chapter 3"}
	Page2Book2 := Page{Number: 2, Content: "chapter 4"}
	Page1Book3 := Page{Number: 1, Content: "chapter 5"}
	Page2Book3 := Page{Number: 2, Content: "chapter 6"}
	Page1Book4 := Page{Number: 1, Content: "chapter 7"}
	Page2Book4 := Page{Number: 2, Content: "chapter 8"}
	pagesDataBook1 := []Page{Page1Book1, Page2Book1}
	pagesDataBook2 := []Page{Page1Book2, Page2Book2}
	pagesDataBook3 := []Page{Page1Book3, Page2Book3}
	pagesDataBook4 := []Page{Page1Book4, Page2Book4}
	bookdata := []Book{}
	bookdata = append(bookdata, Book{Author: "Vishwas", Pages: pagesDataBook1, Name: "DSA"})
	bookdata = append(bookdata, Book{Author: "Mallikarjuna", Pages: pagesDataBook2, Name: "Go"})
	bookdata = append(bookdata, Book{Author: "Manav", Pages: pagesDataBook3, Name: "Python"})
	bookdata = append(bookdata, Book{Author: "Bhora", Pages: pagesDataBook4, Name: "Java"})
	library := Library{bookdata}

	fmt.Println(library.books)

	router := gin.Default()

	router.GET("/:author/:name", func(c *gin.Context) {
		bookauthour := c.Param("author")
		bookname := c.Param("name")
		ReturnedPageData := library.getPages(bookauthour, bookname)
		c.JSON(http.StatusOK, ReturnedPageData)
	})
	router.Run(":8000")
}

func (L1 Library) getPages(authourname string, bookname string) []Page {
	for i := 0; i < len(L1.books); i++ {
		if L1.books[i].Author == authourname && L1.books[i].Name == bookname {
			return L1.books[i].Pages
		}
	}
	return []Page{}
}
