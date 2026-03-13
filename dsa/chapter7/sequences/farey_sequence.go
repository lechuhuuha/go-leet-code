package sequences

import "fmt"

type Fraction struct {
	Numerator   int
	Denominator int
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}

func g(l Fraction, r Fraction, num int) {
	frac := Fraction{
		Numerator:   l.Numerator + r.Numerator,
		Denominator: l.Denominator + r.Denominator,
	}

	if frac.Denominator < num {
		g(l, frac, num)
		fmt.Println(frac)
		g(frac, r, num)
	}
}
