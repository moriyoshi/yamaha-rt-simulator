package simulator

import "context"

var CmdConsoleColumns = &CommandSpec{
	[]Token{TConsole, TColumns}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		return nil
	},
}

var CmdNoConsoleColumns = &CommandSpec{
	[]Token{TNo, TConsole, TColumns}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		return nil
	},
}
