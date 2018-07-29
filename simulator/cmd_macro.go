package simulator

import "context"

var CmdMacro = &CommandSpec{
	[]Token{TMacro}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMacroNameEom = &CommandSpec{
	[]Token{TNo, TMacro}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

