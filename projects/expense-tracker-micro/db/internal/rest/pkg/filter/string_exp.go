package filter

type StringExp struct {
	IsNullMatch *bool    `json:"isNull,omitempty"`
	EqMatch     *string  `json:"eq,omitempty"`
	NeqMatch    *string  `json:"neq,omitempty"`
	InMatch     []string `json:"in,omitempty"`
	NinMatch    []string `json:"nin,omitempty"`
	LikeMatch   *string  `json:"like,omitempty"`
	NlikeMatch  *string  `json:"nlike,omitempty"`
	IlikeMatch  *string  `json:"ilike,omitempty"`
	NilikeMatch *string  `json:"nilike,omitempty"`
}

func (exp *StringExp) IsNull() *bool {
	return exp.IsNullMatch
}

func (exp *StringExp) Eq() *string {
	return exp.EqMatch
}

func (exp *StringExp) Neq() *string {
	return exp.NeqMatch
}

func (exp *StringExp) In() []string {
	return exp.InMatch
}

func (exp *StringExp) Nin() []string {
	return exp.NinMatch
}

func (exp *StringExp) Like() *string {
	return exp.LikeMatch
}

func (exp *StringExp) Nlike() *string {
	return exp.NlikeMatch
}

func (exp *StringExp) Ilike() *string {
	return exp.IlikeMatch
}

func (exp *StringExp) Nilike() *string {
	return exp.NilikeMatch
}
