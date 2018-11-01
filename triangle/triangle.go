package triangle

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	res := getNaT(a, b, c)
	if !res {
		k = NaT
		return k
	}
	res = getEqu(a, b, c)
	if res && k != NaT {
		k = Equ
	}

	res = getIso(a, b, c)
	if res && k != Equ {
		k = Iso
	}

	res = getSca(a, b, c)
	if res && k != Equ {
		k = Sca
	}

	return k
}

func getNaT(a, b, c float64) bool {
	if a > 0 && b > 0 && c > 0 {
		if a+b > c {
			return true
		}
		return true
	}
	return false
}

func getEqu(a, b, c float64) bool {
	if a == b && b == c && a == c {
		return true
	}
	return false
}

func getIso(a, b, c float64) bool {
	if a == b || b == c || a == c {
		return true
	}
	return false
}

func getSca(a, b, c float64) bool {
	if a != b && b != c && a != c || 2*a == b+c {
		return true
	}
	return false
}
