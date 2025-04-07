package main

import "fmt"

func main() {
	bookmarks := map[string]string{}
	fmt.Println("Hello, this is Marky.")
Menu:
	for {
		option := getMenu()
		switch option {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			bookmarks = addBookmark(bookmarks)
		case 3:
			bookmarks = deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var option int
	fmt.Println("Select an option")
	fmt.Println("1. View bookmarks")
	fmt.Println("2. Add a bookmark")
	fmt.Println("3. Delete a bookmark")
	fmt.Println("4. Exit")
	fmt.Scan(&option)
	return option
}

func printBookmarks(bookmarks map[string]string) {
	if len(bookmarks) == 0 {
		fmt.Println("There are no bookmarks yet")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ": ", value)

	}
}

func addBookmark(bookmarks map[string]string) map[string]string {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Enter title: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Enter link: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
	return bookmarks
}

func deleteBookmark(bookmarks map[string]string) map[string]string {
	var bookmarkKetToDelete string
	fmt.Print("Enter title: ")
	fmt.Scan(&bookmarkKetToDelete)
	delete(bookmarks, bookmarkKetToDelete)
	return bookmarks
}
