package simulator

import "context"

var CmdCloudVpnService = &CommandSpec{
	[]Token{TCloud, TVpn, TService}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoCloudVpnService = &CommandSpec{
	[]Token{TNo, TCloud, TVpn, TService}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

