package main

import (
	"fmt"
	" hery-ciaputra/demo-gin/db"
	" hery-ciaputra/demo-gin/server"
)

//func myMiddleware(text string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		name, _ := c.Get("user")
//		fmt.Println("MyMiddleware: ", text, name)
//	}
//}

func main() {
	err := db.Connect()
	if err != nil {
		fmt.Println("failed to connect to DB")
	}

	server.Init()
}
