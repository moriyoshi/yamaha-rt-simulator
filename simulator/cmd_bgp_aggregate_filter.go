package simulator

import "context"

var CmdBgpAggregateFilter = &CommandSpec{
	[]Token{TBgp, TAggregate, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpAggregateFilter = &CommandSpec{
	[]Token{TNo, TBgp, TAggregate, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

