package simulator

import "context"

var CmdWanAuthMyname = &CommandSpec{
	[]Token{TWan, TAuth, TMyname}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanAuthMyname = &CommandSpec{
	[]Token{TNo, TWan, TAuth, TMyname}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

