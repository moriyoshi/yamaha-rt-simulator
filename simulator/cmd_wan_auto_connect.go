package simulator

import "context"

var CmdWanAutoConnect = &CommandSpec{
	[]Token{TWan, TAuto, TConnect}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanAutoConnect = &CommandSpec{
	[]Token{TNo, TWan, TAuto, TConnect}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

