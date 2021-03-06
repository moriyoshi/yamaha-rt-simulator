package simulator

import "context"

var CmdHelp = &CommandSpec{
	[]Token{THelp}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

