package simulator

import "context"

var CmdConsoleLines = &CommandSpec{
	[]Token{TConsole, TLines, &EitherToken{TInfinity, TSomething}}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		return nil
	},
}

var CmdNoConsoleLines = &CommandSpec{
	[]Token{TNo, TConsole, TLines, &EitherToken{TInfinity, TSomething}}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		return nil
	},
}
