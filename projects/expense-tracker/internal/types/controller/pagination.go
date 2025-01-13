package controller

type Pagination interface {
	GetLimit() int
	GetOffset() int
}
