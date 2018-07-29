package simulator

import "context"

var CmdWolSend = &CommandSpec{
	[]Token{TWol, TSend}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
