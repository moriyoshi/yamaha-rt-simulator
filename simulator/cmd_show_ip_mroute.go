package simulator

import "context"

var CmdShowIpMroute = &CommandSpec{
	[]Token{TShow, TIp, TMroute}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

