package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	ID string     	`json:"id"`
	Title string   	`json:"title"`
	Author string	`json:"author"`
	Quantity int	`json: "quantity"`

}



// JSON - Javascript object notation