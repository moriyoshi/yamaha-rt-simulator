package simulator

import "context"

var CmdL2TpKeepaliveUse = &CommandSpec{
	[]Token{TL2Tp, TKeepalive, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoL2TpKeepaliveUse = &CommandSpec{
	[]Token{TNo, TL2Tp, TKeepalive, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

