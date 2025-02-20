package filter

type ComparableExp[T any] struct {
	IsNullMatch *bool `json:"isNull,omitempty"`
	EqMatch     *T    `json:"eq,omitempty"`
	NeqMatch    *T    `json:"neq,omitempty"`
	InMatch     []T   `json:"in,omitempty"`
	NinMatch    []T   `json:"nin,omitempty"`
	GtMatch     *T    `json:"gt,omitempty"`
	GteMatch    *T    `json:"gte,omitempty"`
	LtMatch     *T    `json:"lt,omitempty"`
	LteMatch    *T    `json:"lte,omitempty"`
}

func (exp *ComparableExp[T]) IsNull() *bool {
	return exp.IsNullMatch
}

func (exp *ComparableExp[T]) Eq() *T {
	return exp.EqMatch
}

func (exp *ComparableExp[T]) Neq() *T {
	return exp.NeqMatch
}

func (exp *ComparableExp[T]) In() []T {
	return exp.InMatch
}

func (exp *ComparableExp[T]) Nin() []T {
	return exp.NinMatch
}

func (exp *ComparableExp[T]) Gt() *T {
	return exp.GtMatch
}

func (exp *ComparableExp[T]) Gte() *T {
	return exp.GteMatch
}

func (exp *ComparableExp[T]) Lt() *T {
	return exp.LtMatch
}

func (exp *ComparableExp[T]) Lte() *T {
	return exp.LteMatch
}
