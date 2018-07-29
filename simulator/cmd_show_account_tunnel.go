package simulator

import "context"

var CmdShowAccountTunnel = &CommandSpec{
	[]Token{TShow, TAccount, TTunnel}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

