package simulator

import "context"

var CmdMakeDirectory = &CommandSpec{
	[]Token{TMake, TDirectory}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

