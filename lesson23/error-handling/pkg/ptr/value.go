package ptr

func ToPtr[V any](v V) *V {
	return &v
}

func ToValue[V any](p *V) V {
	var v V
	if p != nil {
		v = *p
	}
	return v
}
