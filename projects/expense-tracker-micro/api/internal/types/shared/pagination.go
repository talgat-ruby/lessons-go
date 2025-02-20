package shared

type PaginationOffset interface {
	GetLimit() int
	GetOffset() int
}
