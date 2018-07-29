package simulator

import "context"

var CmdWanAlwaysOn = &CommandSpec{
	[]Token{TWan, TAlwaysOn}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanAlwaysOn = &CommandSpec{
	[]Token{TNo, TWan, TAlwaysOn}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

