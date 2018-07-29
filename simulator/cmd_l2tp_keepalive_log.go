package simulator

import "context"

var CmdL2TpKeepaliveLog = &CommandSpec{
	[]Token{TL2Tp, TKeepalive, TLog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoL2TpKeepaliveLog = &CommandSpec{
	[]Token{TNo, TL2Tp, TKeepalive, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

