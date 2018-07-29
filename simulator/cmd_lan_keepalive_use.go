package simulator

import "context"

var CmdLanKeepaliveUse = &CommandSpec{
	[]Token{TLan, TKeepalive, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanKeepaliveUse = &CommandSpec{
	[]Token{TNo, TLan, TKeepalive, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
