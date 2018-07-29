package simulator

import "context"

var CmdCloudVpnSetGo = &CommandSpec{
	[]Token{TCloud, TVpn, TSet, TGo}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

