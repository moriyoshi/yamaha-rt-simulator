package simulator

import "context"

var CmdIpKeepalive = &CommandSpec{
	[]Token{TIp, TKeepalive}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpKeepalive = &CommandSpec{
	[]Token{TNo, TIp, TKeepalive}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

