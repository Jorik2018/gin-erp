package repository

import (
	"container/list"
	_ "database/sql"
	"reflect"
	//"log"

	//util "github.com/Jorik2018/db-helper/mysql"
	"github.com/Jorik2018/gin-erp/models"
	"github.com/Jorik2018/gin-erp/db"
)

//BookRepo - Book repository
type BookRepo struct {
	Name string
}

//GetBookRepository - returns book repository
func GetBookRepository() BookRepo {
	return BookRepo{Name: "tbl_book"}
}

// func toList(slice interface{}) *list.List {
//     sliceValue := slice.([]interface{})
//     bookList := list.New()
//     for _, item := range sliceValue {
//         bookList.PushBack(item)
//     }
//     return bookList
// }

func toList(slice interface{}) *list.List {
    sliceValue := reflect.ValueOf(slice)

    // Ensure sliceValue is a slice
    if sliceValue.Kind() != reflect.Slice {
        panic("toList: input is not a slice")
    }

    bookList := list.New()
    for i := 0; i < sliceValue.Len(); i++ {
        bookList.PushBack(sliceValue.Index(i).Interface())
    }
    return bookList
}

func (repo BookRepo) Select() (*list.List, error) {
	var books [] models.Book
	var result = db.ORM.Find(&books)
	return toList(books), result.Error
	// reader := func(rows *sql.Rows, collection *list.List) {
	// 	var book dao.Book
	// 	err := rows.Scan(&book.BookID, &book.BookName, &book.Author)
	// 	collection.PushBack(book)
	// 	log.Fatal(err)
	// }
	// return util.ExecuteReader(DbConnection, "select book_id, book_name, book_author from tbl_book", reader)
}

func (repo BookRepo) Find(id string) (models.Book, error) {
    var book models.Book
    if err := db.ORM.Where("id = ?", id).First(&book).Error; err != nil {
        return book, err
    }
    return book, nil
}

//Insert - Insert books to db
func (repo BookRepo) Insert(book *models.Book) (int64, error) {
  	result := db.ORM.Create(&book)
	return 0, result.Error
	// book := doc.(dao.Book)
	// return util.ExecuteInsert(DbConnection, "insert into tbl_book(book_name, book_author) values (?,?)", book.BookName, book.Author)
}

//Update - Update books
func (repo BookRepo) Update(doc interface{}) (int64, error) {
	return 0, nil
	// book := doc.(dao.Book)
	// return util.ExecuteUpdateDelete(DbConnection, "update tbl_book set book_name=?, book_author=? where book_id=?", book.BookName, book.Author, book.BookID)
}

//Remove - Delete books from db
func (repo BookRepo) Remove(doc interface{}) (int64, error) {
	return 0, nil
	// book := doc.(dao.Book)
	// return util.ExecuteUpdateDelete(DbConnection, "delete from tbl_book where book_id=?", book.BookID)
}
