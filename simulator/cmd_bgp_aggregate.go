package simulator

import "context"

var CmdBgpAggregate = &CommandSpec{
	[]Token{TBgp, TAggregate}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpAggregate = &CommandSpec{
	[]Token{TNo, TBgp, TAggregate}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

