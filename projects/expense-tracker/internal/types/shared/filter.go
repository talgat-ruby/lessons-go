package shared

import "time"

type Filter[F any] interface {
	GetFilter() F
}

type FilterItem[F any] interface {
	FilterAnd() []F
	FilterOr() []F
}

type IsNullExp interface {
	IsNull() *bool
}

type EqExp[T any] interface {
	Eq() *T
	Neq() *T
}

type InExp[T any] interface {
	In() []T
	Nin() []T
}

type LikeExp[T any] interface {
	Like() *T
	Nlike() *T
	Ilike() *T
	Nilike() *T
}

type QuantifyExp[T any] interface {
	Gt() *T
	Gte() *T
	Lt() *T
	Lte() *T
}

type IdenticalExp[T any] interface {
	IsNullExp
	EqExp[T]
	InExp[T]
}

type ComparableExp[T any] interface {
	IdenticalExp[T]
	QuantifyExp[T]
}

type IDExp interface {
	IdenticalExp[string]
}

type BoolExp interface {
	IsNullExp
	EqExp[bool]
}

type StringExp interface {
	IdenticalExp[string]
	LikeExp[string]
}

type Int32Exp interface {
	ComparableExp[int32]
}

type Int64Exp interface {
	ComparableExp[int64]
}

type Float32Exp interface {
	ComparableExp[float32]
}

type Float64Exp interface {
	ComparableExp[float64]
}

type TimeExp interface {
	ComparableExp[time.Time]
}
