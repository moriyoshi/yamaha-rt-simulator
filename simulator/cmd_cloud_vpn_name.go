package simulator

import "context"

var CmdCloudVpnName = &CommandSpec{
	[]Token{TCloud, TVpn, TName}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoCloudVpnName = &CommandSpec{
	[]Token{TNo, TCloud, TVpn, TName}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

