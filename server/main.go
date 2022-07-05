package main

import (
	"fmt"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct{
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Quantity int	`json:"quantity"`
}

var books = []book{
		{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
		{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
		{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK,books)
}

func booksByID(c *gin.Context){
	id:=c.Param("id")
	book,err:=getBooksById(id)
	fmt.Println(err)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Book Not Found"})
		return
	}
	c.IndentedJSON(http.StatusOK,book)
}

func getBooksById(id string)(*book,error){
	for i,b:=range books{
		if b.ID==id{
			return &books[i],nil
		}
	}
	return  nil,errors.New("Books Not Found")
}

func  createBooks(c *gin.Context){
	var newbook book

	if err:=c.BindJSON(&newbook);err!=nil{
		return
	}
	books=append(books,newbook)
	c.IndentedJSON(http.StatusCreated,gin.H{"message":"Book created"})
}

func updateBook(c *gin.Context){
	id,ok:=c.GetQuery("id")
	title,_:=c.GetQuery("title")

	if !ok{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Missing id parameter"})
		return
	}
	
	book,err:=getBooksById(id)
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Book Not Found"})
		return
	}
    book.Title=title
	c.IndentedJSON(http.StatusOK,book)
}


func main(){
	router:=gin.Default()
	router.GET("/book",getBooks)
	router.POST("/book",createBooks)
	router.GET("/book/:id",booksByID)
	router.PATCH("/book",updateBook)
	router.Run("localhost:8080")
}