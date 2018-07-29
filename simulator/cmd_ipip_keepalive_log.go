package simulator

import "context"

var CmdIpipKeepaliveLog = &CommandSpec{
	[]Token{TIpip, TKeepalive, TLog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpipKeepaliveLog = &CommandSpec{
	[]Token{TNo, TIpip, TKeepalive, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

