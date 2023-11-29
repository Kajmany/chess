package uci

//My patch: This whole file is mine (you can tell) but modeled substantially off the author's work that already existed

// notnil's position command would require me to start a whole chess game with the parent library OR do a lot of manual
// parsing in order to get chess.Position, but I really just want to pass my FEN to the engine!
type CmdDumbPosition struct {
	//Forsyth-Edwards Notation representation of chess board & game state
	FEN string
}

// This just cuts out the extra logic in uci.CmdPosition and echoes the string as a UCI command for the uci.Engine
func (cmd CmdDumbPosition) String() string {
	return "position fen " + cmd.FEN
}

// Required to satisfy interface. As with uci.CmdPosition, there's no possible UCI error state because the engine will
// try to ignore a bad position. !!Stockfish won't try hard and will crash often!!
func (cmd CmdDumbPosition) ProcessResponse(e *Engine) error {
	return nil
}
