package minivm

type Value interface {
	Value() interface{}
	tobool() bool
	add(Value) Value
	sub(Value) Value
	mul(Value) Value
	div(Value) Value
}

type VBool struct {
	value bool
}

func (v VBool) Value() interface{} {
	return v.value
}

func (v VBool) tobool() bool {
	return v.value
}

func (rhs VBool) add(lhs Value) Value {
	panic("you cannot add boolean")
}

func (rhs VBool) sub(lhs Value) Value {
	panic("you cannot subtract boolean")
}

func (rhs VBool) mul(lhs Value) Value {
	panic("you cannot multiply by boolean")
}

func (rhs VBool) div(lhs Value) Value {
	panic("you cannot divide by boolean")
}

type VInt struct {
	value int64
}

func (v VInt) Value() interface{} {
	return v.value
}

func (v VInt) tobool() bool {
	panic("you cannot use int for boolean")
}

func (rhs VInt) add(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VInt{value: lhs.value + rhs.value}
	case VFloat:
		return VFloat{value: lhs.value + float64(rhs.value)}
	default:
		panic("invalid value type for add")
	}
}

func (rhs VInt) sub(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VInt{value: lhs.value - rhs.value}
	case VFloat:
		return VFloat{value: lhs.value - float64(rhs.value)}
	default:
		panic("invalid value type for sub")
	}
}

func (rhs VInt) mul(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VInt{value: lhs.value * rhs.value}
	case VFloat:
		return VFloat{value: lhs.value * float64(rhs.value)}
	default:
		panic("invalid value type for mul")
	}
}

func (rhs VInt) div(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VInt{value: lhs.value / rhs.value}
	case VFloat:
		return VFloat{value: lhs.value / float64(rhs.value)}
	default:
		panic("invalid value type for div")
	}
}

type VFloat struct {
	value float64
}

func (v VFloat) Value() interface{} {
	return v.value
}

func (v VFloat) tobool() bool {
	panic("you cannot use float for boolean")
}

func (rhs VFloat) add(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VFloat{value: float64(lhs.value) + rhs.value}
	case VFloat:
		return VFloat{value: lhs.value + rhs.value}
	default:
		panic("invalid value type for add")
	}
}

func (rhs VFloat) sub(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VFloat{value: float64(lhs.value) - rhs.value}
	case VFloat:
		return VFloat{value: lhs.value - rhs.value}
	default:
		panic("invalid value type for sub")
	}
}

func (rhs VFloat) mul(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VFloat{value: float64(lhs.value) * rhs.value}
	case VFloat:
		return VFloat{value: lhs.value * rhs.value}
	default:
		panic("invalid value type for mul")
	}
}

func (rhs VFloat) div(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VFloat{value: float64(lhs.value) / rhs.value}
	case VFloat:
		return VFloat{value: lhs.value / rhs.value}
	default:
		panic("invalid value type for div")
	}
}
