package model

import (
	"sync"
	"errors"
)

var (
	lock sync.Mutex
	ErrStockNotEnough = errors.New("stock is not enough")

)

type Book struct {
	Name string
	Total int
	Author string
	CreateTime string
}

//创建书
func CreateBook(name string, total int, author string, createTime string) *Book  {
	return  &Book{
		Name:name,
		Total:total,
		Author:author,
		CreateTime:createTime,
	}
}

//判断是否可借
func (b *Book) CanBorrow(c int) bool {
	return b.Total >= c
}

//借书
func (b *Book) Borrow(c int) (err error) {
	if !b.CanBorrow(c) {
		err = ErrStockNotEnough
		return
	}

	lock.Lock()
	b.Total -= c
	lock.Unlock()
	return
}
//还书
func (b *Book) Back(c int){
	lock.Lock()
	b.Total += c
	lock.Unlock()
	return
}