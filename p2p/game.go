package p2p

type GameVariant uint8

const (
	TexasHoldem GameVariant = iota
	Other
)

func (gv GameVariant) String() string {
	switch gv {
	case TexasHoldem:
		return "TEXAS HOLD'EM"
	case Other:
		return "OTHER"
	default:
		return "UNKNOWN GAME VARIANT"
	}
}
