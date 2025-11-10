package consensus

// Placeholder: in devnet we operate with no consensus.
// Future options: PoS, HotStuff variants, PQ-secure signatures, etc.
type Engine struct{}

func NewNone() *Engine { return &Engine{} }
