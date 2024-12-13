package vec

type Line struct {
	a Vec
	b Vec
}

func MakeLine(a Vec, b Vec) Line {
	return Line {
		a: a,
		b: b,
	}
}

func (l *Line) OnTheLine(v Vec) bool {
	if v == l.a || v == l.b {
		return true
	}
	aToB := l.a.Sub(l.b)
	aToBStep := aToB.NormalizeToInt()
	checkVec := l.a
	for checkVec != l.b {
		checkVec = checkVec.Add(aToBStep)
		if checkVec == v {
			return true
		}
	}
	return false
}
