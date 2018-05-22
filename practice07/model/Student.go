package model

import (
	"sync"
	"errors"
)

var (
	sLock sync.Mutex
	ErrBookNotExist = errors.New("the book is not exist")
	ErrBookNumWrong = errors.New("the deleted book num is wrong")
)

type Student struct {
	Name string
	Grade string
	Id string
	Sex string
	//指针存的结构小
	books []*BorrowItem
}

type BorrowItem struct {
	Book *Book
	num int
}

func CreateStudent(name, grade, id, sex string) *Student {
	return &Student{
		Name:name,
		Grade:grade,
		Id:id,
		Sex:sex,
	}
}

func (s *Student) AddBook(b *BorrowItem) {
	s.books = append(s.books, b)
}

func (s *Student) DelBook(b *BorrowItem) (err error) {
	for i := 0; i < len(s.books); i++ {
		if s.books[i].Book.Name == b.Book.Name {
			if s.books[i].num == b.num {
				front := s.books[:i]
				right := s.books[i+1:]
				front = append(front, right...)
				s.books = front
				return
			}else if s.books[i].num > b.num {
				s.books[i].num -= b.num
				return
			}else {
				err = ErrBookNumWrong
			}
		}
	}
	if err == nil {
		err = ErrBookNotExist
	}
	return
}

//返回借书列表
func (s *Student) GetBookList() []*BorrowItem {
	return s.books
}