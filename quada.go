package piscine

import "github.com/01-edu/z01"

func Quad(x, y int, tlc, trc, blc, brc, vl, hl rune) {
	if x <= 0 || y <= 0 {
		return
	}
	for h := 1; h <= y; h++ {
		isTopRow := h == 1
		isBottomRow := h == y

		for w := 1; w <= x; w++ {
			isLeftCol := w == 1
			isRightCol := w == x

			isTopLeft := isTopRow && isLeftCol
			isTopRight := isTopRow && isRightCol
			isBottomLeft := isBottomRow && isLeftCol
			isBottomRight := isBottomRow && isRightCol

			switch {
			case isTopLeft:
				z01.PrintRune(tlc)
			case isTopRight:
				z01.PrintRune(trc)
			case isBottomLeft:
				z01.PrintRune(blc)
			case isBottomRight:
				z01.PrintRune(brc)
			case isTopRow || isBottomRow:
				z01.PrintRune(hl)
			case isLeftCol || isRightCol:
				z01.PrintRune(vl)
			default:
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadA(x, y int) {
	Quad(x, y, 'o', 'o', 'o', 'o', '|', '-')
}
