package simulator

import "context"

var CmdLanKeepaliveLog = &CommandSpec{
	[]Token{TLan, TKeepalive, TLog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanKeepaliveLog = &CommandSpec{
	[]Token{TNo, TLan, TKeepalive, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

