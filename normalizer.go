package janorm

import (
	"strings"

	"github.com/koron-go/trietree"
)

type builder struct {
	dt *trietree.DTree
	ws []string
}

func newBuilder(n int) *builder {
	return &builder{
		dt: &trietree.DTree{},
		ws: make([]string, 0, n),
	}
}

func (b *builder) put(from, to string) {
	n := b.dt.Put(from) - 1
	if n < len(b.ws) {
		b.ws[n] = to
		return
	}
	b.ws = append(b.ws, to)
}

func (b *builder) putEach(from string, fromUnit int, to string, toUnit int) {
	var (
		fromRune = []rune(from)
		toRune   = []rune(to)
	)
	if len(fromRune)*toUnit != len(toRune)*fromUnit {
		panic("rune length mismatch")
	}
	n := len(fromRune) / fromUnit
	x, y := 0, 0
	for i := 0; i < n; i++ {
		b.put(string(fromRune[x:x+fromUnit]), string(toRune[y:y+toUnit]))
		x += fromUnit
		y += toUnit
	}
}

func (b *builder) putMap(to string, froms ...string) {
	for _, from := range froms {
		b.put(from, to)
	}
}

func (b *builder) normalizer() *normalizer {
	st := trietree.Freeze(b.dt)
	return &normalizer{
		st: st,
		ws: b.ws,
	}
}

type normalizer struct {
	st *trietree.STree
	ws []string
}

func (n *normalizer) normalize(s string) string {
	sc := newScanner(s)
	n.st.Scan(s, sc)
	return sc.finish(n.ws)
}

type scanner struct {
	rs []rune

	ids   []int
	lasts []int

	ridx int
}

func newScanner(s string) *scanner {
	rs := []rune(s)
	lasts := make([]int, len(rs))
	for i := range lasts {
		lasts[i] = i
	}
	return &scanner{
		rs:    rs,
		ids:   make([]int, len(rs)),
		lasts: lasts,
	}
}

func (s *scanner) finish(ws []string) string {
	b := &strings.Builder{}
	for i := 0; i < len(s.rs); i++ {
		id := s.ids[i]
		if id == 0 {
			b.WriteRune(s.rs[i])
		} else {
			b.WriteString(ws[id-1])
		}
		i = s.lasts[i]
	}
	return b.String()
}

func (s *scanner) ScanReport(ev trietree.ScanEvent) {
	ridx := s.ridx
	s.ridx++
	if len(ev.Nodes) == 0 {
		return
	}
	for _, n := range ev.Nodes {
		start := ridx - n.Level + 1
		last := ridx
		if s.ids[start] == 0 || last > s.lasts[start] {
			s.ids[start] = n.ID
			s.lasts[start] = last
		}
	}
}
