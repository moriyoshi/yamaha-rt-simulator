package simulator

import "context"

var CmdShowIpRoute = &CommandSpec{
	[]Token{TShow, TIp, TRoute}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
