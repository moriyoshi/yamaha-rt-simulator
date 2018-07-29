package simulator

import "context"

var CmdWanAccessPointName = &CommandSpec{
	[]Token{TWan, TAccessPoint, TName}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanAccessPointName = &CommandSpec{
	[]Token{TNo, TWan, TAccessPoint, TName}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

