package simulator

import "context"

var CmdRdate = &CommandSpec{
	[]Token{TRdate}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

