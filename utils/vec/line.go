package vec

type Line struct {
	A Vec
	B Vec
}

func MakeLine(a Vec, b Vec) Line {
	if a == b {
		panic("a and b are the name vec")
	}
	return Line {
		A: a,
		B: b,
	}
}

func (l *Line) OnTheLine(v Vec) bool {
	if v == l.A || v == l.B {
		return true
	}
	aToB := l.AToB()
	aToBStep := aToB.NormalizeToInt()
	checkVec := l.A
	for checkVec != l.B {
		checkVec = checkVec.Add(aToBStep)
		if checkVec == v {
			return true
		}
	}
	return false
}

func (l *Line) BToA() Vec {
	return l.A.Sub(l.B)
}

func (l *Line) AToB() Vec {
	return l.B.Sub(l.A)
}
