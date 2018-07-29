package simulator

import "context"

var CmdShowMacro = &CommandSpec{
	[]Token{TShow, TMacro}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

