package simulator

import "context"

var CmdCloudVpnParameter = &CommandSpec{
	[]Token{TCloud, TVpn, TParameter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoCloudVpnParameter = &CommandSpec{
	[]Token{TNo, TCloud, TVpn, TParameter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

