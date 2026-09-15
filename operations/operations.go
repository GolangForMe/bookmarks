package operations

import "fmt"

func InputBookmark() string {
	var value string
	fmt.Println("Введите название закладки:")
	fmt.Scanln(&value)
	return value
}

func InputKey() int {
	var value int
	fmt.Println("Введите номер закладки:")
	fmt.Scanln(&value)
	return value
}

func GetAllBookmarks(bookmarks map[int]string) {
	if len(bookmarks) == 0 {
		fmt.Println("Список закладок пустой!")
	} else {
		fmt.Println("Список закладок:")
		for key, value := range bookmarks {
			fmt.Println(key, value)
		}
	}
}

func AddNewBookmark(bookmarks map[int]string, bookmark string) {
	maxKey := 0
	for key := range bookmarks {
		if key > maxKey {
			maxKey = key
		}
	}
	bookmarks[maxKey+1] = bookmark
	fmt.Println("Закладка успешно создана!")
}

func DeleteBookmark(bookmarks map[int]string, bookmarkKey int) {
	msg := fmt.Sprintf("Закладка с ключом %d не найдена", bookmarkKey)

	if _, ok := bookmarks[bookmarkKey]; ok {
		delete(bookmarks, bookmarkKey)
		msg = fmt.Sprintf("Закладка с ключом %d успешно удалена!", bookmarkKey)
	}

	fmt.Println(msg)
}
