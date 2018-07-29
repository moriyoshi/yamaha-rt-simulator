package simulator

import "context"

var CmdLeasedKeepaliveDown = &CommandSpec{
	[]Token{TLeased, TKeepalive, TDown}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLeasedKeepaliveDown = &CommandSpec{
	[]Token{TNo, TLeased, TKeepalive, TDown}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

