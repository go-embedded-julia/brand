# go-embedded-julia — brand assets

Official logos for the **go-embedded-julia** organization.

<img src="avatar/go-embedded-julia.png" width="160" alt="go-embedded-julia">

The mark is Julia's subtype operator, `<:` ("is a subtype of"), the relation multiple dispatch is built on: a chevron and a colon. It is a glyph of the language's syntax, and it does not copy the Julia logo. It is written by `mark/main.go`, which centres the glyph on the badge from a few constants; nothing is placed by hand. Its gradient (`#957EF1` to `#4129A3`) sets go-embedded-julia apart from the other go-* organizations, including go-embedded-ruby; its dark end reads on white at 10.1:1.

## Formats

- **SVG** — `svg/color`, `svg/white` (white, transparent), `svg/black` (black, transparent)
- **PNG** 16→1024 px, 3 variants — `png/<variant>/<size>/`
- **JPG** (color, white background) — `jpg/<size>/`
- **ICO** (Windows) — `ico/`
- **ICNS** (macOS) — `icns/`
- **GitHub avatar** 512 px — `avatar/`
- **Social preview** 1280×640 (repo banner) — `social/`

`check/kit.py` counts the whole matrix on every pull request, and the `mark` job runs `mark/main.go` again and fails if `svg/color/go-embedded-julia.svg` differs.

---
*Generated from `svg/color/go-embedded-julia.svg` by the fleet's Go brand pipeline (brandgen) — rounded square badge, line glyph, color / white / black variants.*
