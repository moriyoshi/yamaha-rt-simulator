package simulator

import "context"

var CmdCloudVpnOption = &CommandSpec{
	[]Token{TCloud, TVpn, TOption}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoCloudVpnOption = &CommandSpec{
	[]Token{TNo, TCloud, TVpn, TOption}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

