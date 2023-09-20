package dao

import "finalpoint/model"

func SelectAll() (BookList []model.BookFirstMsg) {
	Db.Table("book_first_msg").Find(&BookList)
	return
}

func Page(begin uint, pagesize int) (bookList []model.BookFirstMsg) {
	Db.Table("book_first_msg").Offset(int(begin)).Limit(pagesize).Find(&bookList)
	return
}
