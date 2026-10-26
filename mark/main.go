// Command mark writes the go-embedded-julia mark, svg/color/go-embedded-julia.svg:
//
//	cd mark && go run . > ../svg/color/go-embedded-julia.svg
//
// go-embedded-julia is a Julia interpreter. The mark is Julia's subtype
// operator, <: (is a subtype of), the relation that multiple dispatch is built
// on: an angle bracket and a colon. It is a glyph of the language's own syntax,
// not the language's logo.
//
// Nothing is placed by hand. The program takes the size of the chevron, the gap
// and the colon from the constants below and centres the whole glyph, ink
// included (the round caps of the chevron and the two dots), on the badge.
// Change a constant and the mark redraws itself.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins. logos/brandgen renders the
// format matrix from the file this writes; the kit workflow runs this again
// and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#957EF1" // the gradient's light stop
	dark  = "#4129A3" // its dark stop

	stroke = 13.0 // the family's line width
	dotR   = 11.0 // radius of a colon dot

	halfH = 54.0 // half the height of the chevron: its arms rise 54 and fall 54
	chevW = 54.0 // the horizontal run of an arm, so each arm is at 45 degrees
	gap   = 30.0 // between the arms' ends and the colon's ink
	colon = 27.0 // the colon's dots sit at the middle line, plus and minus this
)

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

func main() {
	// The ink runs from the chevron tip's cap, x0 - stroke/2, to the colon's
	// right edge, x0 + chevW + gap + 2*dotR. Centre that span on x = 128.
	x0 := (256 + stroke/2 - chevW - gap - 2*dotR) / 2
	// The arms start at the tip and run right: the tip is at the left, so the
	// glyph reads "<".
	tipX, midY := x0, 128.0
	armX := x0 + chevW
	dotX := armX + gap + dotR

	chevron := fmt.Sprintf(`<path d="M%s %sL%s %sL%s %s"/>`,
		num(armX), num(midY-halfH), num(tipX), num(midY), num(armX), num(midY+halfH))
	dots := []string{
		fmt.Sprintf(`<circle cx="%s" cy="%s" r="%s"/>`, num(dotX), num(midY-colon), num(dotR)),
		fmt.Sprintf(`<circle cx="%s" cy="%s" r="%s"/>`, num(dotX), num(midY+colon), num(dotR)),
	}

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-embedded-julia">
  <defs><linearGradient id="ejgrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#ejgrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    %s
  </g>
  <g fill="#FFFFFF">
    %s
  </g>
</svg>
`, light, dark, num(stroke), chevron, strings.Join(dots, "\n    "))
}
