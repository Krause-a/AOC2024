package vec

type Box struct {
	line Line
}

func MakeBox(a Vec, b Vec) Box {
	return Box {
		line: Line {
			A: a,
			B: b,
		},
	}
}

func (b *Box) Contains(v Vec) bool {
	xMin := min(b.line.A.x, b.line.B.x)
	xMax := max(b.line.A.x, b.line.B.x)
	yMin := min(b.line.A.y, b.line.B.y)
	yMax := max(b.line.A.y, b.line.B.y)

	return xMin <= v.x && v.x <= xMax && yMin <= v.y && v.y <= yMax
}
