package simulator

import "context"

var CmdShowStatusTunnel = &CommandSpec{
	[]Token{TShow, TStatus, TTunnel}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
