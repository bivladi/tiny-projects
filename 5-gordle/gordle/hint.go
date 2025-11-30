package gordle

import "strings"

type hint byte
type feedback []hint

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜️"
	case wrongPosition:
		return "🟡"
	case correctPosition:
		return "💚"
	default:
		return "💔"
	}
}

func (f feedback) String() string {
	sb := strings.Builder{}
	for _, h := range f {
		sb.WriteString(h.String())
	}
	return sb.String()
}

func (f feedback) Equal(other feedback) bool {
	if len(f) != len(other) {
		return false
	}
	for idx := range f {
		if f[idx] != other[idx] {
			return false
		}
	}
	return true
}