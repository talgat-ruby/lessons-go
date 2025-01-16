package filter

type IDExp struct {
	IsNullMatch *bool    `json:"isNull,omitempty"`
	EqMatch     *string  `json:"eq,omitempty"`
	NeqMatch    *string  `json:"neq,omitempty"`
	InMatch     []string `json:"in,omitempty"`
	NinMatch    []string `json:"nin,omitempty"`
}

func (exp *IDExp) IsNull() *bool {
	return exp.IsNullMatch
}

func (exp *IDExp) Eq() *string {
	return exp.EqMatch
}

func (exp *IDExp) Neq() *string {
	return exp.NeqMatch
}

func (exp *IDExp) In() []string {
	return exp.InMatch
}

func (exp *IDExp) Nin() []string {
	return exp.NinMatch
}
