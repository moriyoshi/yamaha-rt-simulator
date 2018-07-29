package simulator

import "context"

var CmdShowStatusIpip = &CommandSpec{
	[]Token{TShow, TStatus, TIpip, TTunnel}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
