package vec

type Vec struct {
	x int
	y int
}

var (
	Up = Vec {x:0, y:-1}
	Down = Vec {x:0, y:1}
	Left = Vec {x:-1, y:0}
	Right = Vec {x:1, y:0}
)

func MakeVec(x int, y int) Vec {
	return Vec {
		x: x,
		y: y,
	}
}

func (v *Vec) Add(other Vec) Vec {
	return Vec {
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v *Vec) Invert() Vec {
	return Vec {
		x: -v.x,
		y: -v.y,
	}
}

func (v *Vec) Rotate() Vec {
	return Vec {
		x: -v.y,
		y: v.x,
	}
}

func (v *Vec) Sub(other Vec) Vec {
	return Vec {
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func (v *Vec) NormalizeToInt() Vec {
	x := v.x
	y := v.y
	for y != 0 {
		x, y = y, x%y
	}
	return Vec {
		x: v.x / x,
		y: v.y / x,
	}
}

func (v *Vec) Index() int {
	vec := *v
	if vec == Up {
		return 0
	} else if vec == Down {
		return 1
	} else if vec == Left {
		return 2
	} else if vec == Right {
		return 3
	}
	panic("Invalid Vec converted to index")
}

func FromIndex(index int) Vec {
	if index == 0 {
		return Up
	} else if index == 1 {
		return Down
	} else if index == 2 {
		return Left
	} else if index == 3 {
		return Right
	}
	panic("Invalid index converted to Vec")
}