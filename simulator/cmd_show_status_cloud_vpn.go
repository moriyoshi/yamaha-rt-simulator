package simulator

import "context"

var CmdShowStatusCloudVpn = &CommandSpec{
	[]Token{TShow, TStatus, TCloud, TVpn}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

