package simulator

import "context"

var CmdWanBind = &CommandSpec{
	[]Token{TWan, TBind}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanBind = &CommandSpec{
	[]Token{TNo, TWan, TBind}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
