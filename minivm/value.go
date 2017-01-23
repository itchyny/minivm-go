package minivm

type Value interface {
	Value() interface{}
}

type VInt struct {
	value int64
}

func (v VInt) Value() interface{} {
	return v.value
}

type VFloat struct {
	value float64
}

func (v VFloat) Value() interface{} {
	return v.value
}
