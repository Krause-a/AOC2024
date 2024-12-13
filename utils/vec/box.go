package vec

type Box struct {
	line Line
}

func MakeBox(a Vec, b Vec) Box {
	return Box {
		line: Line {
			a: a,
			b: b,
		},
	}
}

func (b *Box) Contains(v Vec) bool {
	xMin := min(b.line.a.x, b.line.b.x)
	xMax := max(b.line.a.x, b.line.b.x)
	yMin := min(b.line.a.y, b.line.b.y)
	yMax := max(b.line.a.y, b.line.b.y)

	return xMin <= v.x && v.x <= xMax && yMin <= v.y && v.y <= yMax
}
