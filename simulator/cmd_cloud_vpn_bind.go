package simulator

import "context"

var CmdCloudVpnBind = &CommandSpec{
	[]Token{TCloud, TVpn, TBind}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoCloudVpnBind = &CommandSpec{
	[]Token{TNo, TCloud, TVpn, TBind}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

