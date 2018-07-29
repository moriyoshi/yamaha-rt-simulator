package simulator

import "context"

var CmdRename = &CommandSpec{
	[]Token{TRename}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

