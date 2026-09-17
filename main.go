package main

import (
	"bookmarks/helpers"
	"bookmarks/operations"
	"fmt"
)

func main() {
	bookmarks := make(operations.BookmarkMap)
	helpers.Greeting()

outerLoop:
	for {
		choice := helpers.UserInput()

		switch choice {
		case 1:
			operations.GetAllBookmarks(bookmarks)
		case 2:
			mark := operations.InputBookmark()
			operations.AddNewBookmark(bookmarks, mark)
		case 3:
			key := operations.InputKey()
			operations.DeleteBookmark(bookmarks, key)
		case 4:
			fmt.Println("До встречи!")
			break outerLoop
		}
	}
}
