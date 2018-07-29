package simulator

import "context"

var CmdIpsecTunnel = &CommandSpec{
	[]Token{TIpsec, TTunnel}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecTunnel = &CommandSpec{
	[]Token{TNo, TIpsec, TTunnel}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

