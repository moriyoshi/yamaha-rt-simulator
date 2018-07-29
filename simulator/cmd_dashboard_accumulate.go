package simulator

import "context"

var CmdDashboardAccumulate = &CommandSpec{
	[]Token{TDashboard, TAccumulate}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDashboardAccumulate = &CommandSpec{
	[]Token{TNo, TDashboard, TAccumulate}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

