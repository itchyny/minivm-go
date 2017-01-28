package minivm

type Value interface {
	Value() interface{}
	tobool() bool
	plus() Value
	minus() Value
	not() Value
	add(Value) Value
	sub(Value) Value
	mul(Value) Value
	div(Value) Value
	gt(Value) Value
	ge(Value) Value
	eq(Value) Value
	neq(Value) Value
	lt(Value) Value
	le(Value) Value
	addi(int64) Value
	subi(int64) Value
	muli(int64) Value
	divi(int64) Value
	gti(int64) Value
	gei(int64) Value
	eqi(int64) Value
	neqi(int64) Value
	lti(int64) Value
	lei(int64) Value
}

type VFunc struct {
	pc   int
	vars int
}

func (v VFunc) Value() interface{} {
	return v.pc
}

func (v VFunc) tobool() bool {
	panic("you cannot use function for boolean")
}

func (v VFunc) plus() Value {
	panic("you cannot use + on function")
}

func (v VFunc) minus() Value {
	panic("you cannot use - on function")
}

func (v VFunc) not() Value {
	panic("you cannot use ! on function")
}

func (rhs VFunc) add(lhs Value) Value {
	panic("you cannot add function")
}

func (rhs VFunc) sub(lhs Value) Value {
	panic("you cannot subtract function")
}

func (rhs VFunc) mul(lhs Value) Value {
	panic("you cannot multiply function")
}

func (rhs VFunc) div(lhs Value) Value {
	panic("you cannot divide function")
}

func (rhs VFunc) gt(lhs Value) Value {
	panic("you cannot use > on function")
}

func (rhs VFunc) ge(lhs Value) Value {
	panic("you cannot use >= on function")
}

func (rhs VFunc) eq(lhs Value) Value {
	panic("you cannot use == on function")
}

func (rhs VFunc) neq(lhs Value) Value {
	panic("you cannot use != on function")
}

func (rhs VFunc) lt(lhs Value) Value {
	panic("you cannot use < on function")
}

func (rhs VFunc) le(lhs Value) Value {
	panic("you cannot use <= on function")
}

func (v VFunc) addi(i int64) Value {
	panic("you cannot add immediate value to function")
}

func (v VFunc) subi(i int64) Value {
	panic("you cannot subtract immediate value from function")
}

func (v VFunc) muli(i int64) Value {
	panic("you cannot multiply function by immediate value")
}

func (v VFunc) divi(i int64) Value {
	panic("you cannot divide function by immediate value")
}

func (v VFunc) gti(i int64) Value {
	panic("you cannot use > on function")
}

func (v VFunc) gei(i int64) Value {
	panic("you cannot use >= on function")
}

func (v VFunc) eqi(i int64) Value {
	panic("you cannot use == on function")
}

func (v VFunc) neqi(i int64) Value {
	panic("you cannot use != on function")
}

func (v VFunc) lti(i int64) Value {
	panic("you cannot use < on function")
}

func (v VFunc) lei(i int64) Value {
	panic("you cannot use <= on function")
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

func (v VBool) plus() Value {
	panic("you cannot use + on boolean")
}

func (v VBool) minus() Value {
	panic("you cannot use - on boolean")
}

func (v VBool) not() Value {
	return VBool{value: !v.value}
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

func (rhs VBool) gt(lhs Value) Value {
	panic("you cannot use > on boolean")
}

func (rhs VBool) ge(lhs Value) Value {
	panic("you cannot use >= on boolean")
}

func (rhs VBool) eq(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VBool:
		return VBool{value: lhs.value == rhs.value}
	default:
		panic("invalid value type for ==")
	}
}

func (rhs VBool) neq(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VBool:
		return VBool{value: lhs.value != rhs.value}
	default:
		panic("invalid value type for !=")
	}
}

func (rhs VBool) lt(lhs Value) Value {
	panic("you cannot use < on boolean")
}

func (rhs VBool) le(lhs Value) Value {
	panic("you cannot use <= on boolean")
}

func (v VBool) addi(i int64) Value {
	panic("you cannot add immediate value to boolean")
}

func (v VBool) subi(i int64) Value {
	panic("you cannot subtract immediate value from boolean")
}

func (v VBool) muli(i int64) Value {
	panic("you cannot multiply boolean by immediate value")
}

func (v VBool) divi(i int64) Value {
	panic("you cannot divide boolean by immediate value")
}

func (v VBool) gti(i int64) Value {
	panic("you cannot use > on boolean")
}

func (v VBool) gei(i int64) Value {
	panic("you cannot use >= on boolean")
}

func (v VBool) eqi(i int64) Value {
	panic("you cannot use == on boolean")
}

func (v VBool) neqi(i int64) Value {
	panic("you cannot use != on boolean")
}

func (v VBool) lti(i int64) Value {
	panic("you cannot use < on boolean")
}

func (v VBool) lei(i int64) Value {
	panic("you cannot use <= on boolean")
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

func (v VInt) plus() Value {
	return v
}

func (v VInt) minus() Value {
	return VInt{value: -v.value}
}

func (v VInt) not() Value {
	panic("you cannot use ! on int")
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

func (rhs VInt) gt(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VBool{value: lhs.value > rhs.value}
	case VFloat:
		return VBool{value: lhs.value > float64(rhs.value)}
	default:
		panic("invalid value type for >")
	}
}

func (rhs VInt) ge(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VBool{value: lhs.value >= rhs.value}
	case VFloat:
		return VBool{value: lhs.value >= float64(rhs.value)}
	default:
		panic("invalid value type for >=")
	}
}

func (rhs VInt) eq(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VBool{value: lhs.value == rhs.value}
	default:
		panic("invalid value type for ==")
	}
}

func (rhs VInt) neq(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VBool{value: lhs.value != rhs.value}
	default:
		panic("invalid value type for !=")
	}
}

func (rhs VInt) lt(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VBool{value: lhs.value < rhs.value}
	case VFloat:
		return VBool{value: lhs.value < float64(rhs.value)}
	default:
		panic("invalid value type for <")
	}
}

func (rhs VInt) le(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VBool{value: lhs.value <= rhs.value}
	case VFloat:
		return VBool{value: lhs.value <= float64(rhs.value)}
	default:
		panic("invalid value type for <=")
	}
}

func (v VInt) addi(i int64) Value {
	return VInt{value: v.value + i}
}

func (v VInt) subi(i int64) Value {
	return VInt{value: v.value - i}
}

func (v VInt) muli(i int64) Value {
	return VInt{value: v.value * i}
}

func (v VInt) divi(i int64) Value {
	return VInt{value: v.value / i}
}

func (v VInt) gti(i int64) Value {
	return VBool{value: v.value > i}
}

func (v VInt) gei(i int64) Value {
	return VBool{value: v.value >= i}
}

func (v VInt) eqi(i int64) Value {
	return VBool{value: v.value == i}
}

func (v VInt) neqi(i int64) Value {
	return VBool{value: v.value != i}
}

func (v VInt) lti(i int64) Value {
	return VBool{value: v.value < i}
}

func (v VInt) lei(i int64) Value {
	return VBool{value: v.value <= i}
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

func (v VFloat) plus() Value {
	return v
}

func (v VFloat) minus() Value {
	return VFloat{value: -v.value}
}

func (v VFloat) not() Value {
	panic("you cannot use ! on float")
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

func (rhs VFloat) gt(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VBool{value: float64(lhs.value) > rhs.value}
	case VFloat:
		return VBool{value: lhs.value > rhs.value}
	default:
		panic("invalid value type for >")
	}
}

func (rhs VFloat) ge(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VBool{value: float64(lhs.value) >= rhs.value}
	case VFloat:
		return VBool{value: lhs.value >= rhs.value}
	default:
		panic("invalid value type for >=")
	}
}

func (rhs VFloat) eq(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VFloat:
		return VBool{value: lhs.value == rhs.value}
	default:
		panic("invalid value type for ==")
	}
}

func (rhs VFloat) neq(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VFloat:
		return VBool{value: lhs.value != rhs.value}
	default:
		panic("invalid value type for !=")
	}
}

func (rhs VFloat) lt(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VBool{value: float64(lhs.value) < rhs.value}
	case VFloat:
		return VBool{value: lhs.value < rhs.value}
	default:
		panic("invalid value type for <")
	}
}

func (rhs VFloat) le(lhs Value) Value {
	switch lhs := lhs.(type) {
	case VInt:
		return VBool{value: float64(lhs.value) <= rhs.value}
	case VFloat:
		return VBool{value: lhs.value <= rhs.value}
	default:
		panic("invalid value type for <=")
	}
}

func (v VFloat) addi(i int64) Value {
	return VFloat{value: v.value + float64(i)}
}

func (v VFloat) subi(i int64) Value {
	return VFloat{value: v.value - float64(i)}
}

func (v VFloat) muli(i int64) Value {
	return VFloat{value: v.value * float64(i)}
}

func (v VFloat) divi(i int64) Value {
	return VFloat{value: v.value / float64(i)}
}

func (v VFloat) gti(i int64) Value {
	return VBool{value: v.value > float64(i)}
}

func (v VFloat) gei(i int64) Value {
	return VBool{value: v.value >= float64(i)}
}

func (v VFloat) eqi(i int64) Value {
	panic("invalid value type for ==")
}

func (v VFloat) neqi(i int64) Value {
	panic("invalid value type for !=")
}

func (v VFloat) lti(i int64) Value {
	return VBool{value: v.value < float64(i)}
}

func (v VFloat) lei(i int64) Value {
	return VBool{value: v.value <= float64(i)}
}
