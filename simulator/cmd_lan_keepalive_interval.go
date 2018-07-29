package simulator

import "context"

var CmdLanKeepaliveInterval = &CommandSpec{
	[]Token{TLan, TKeepalive, TInterval}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanKeepaliveInterval = &CommandSpec{
	[]Token{TNo, TLan, TKeepalive, TInterval}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

