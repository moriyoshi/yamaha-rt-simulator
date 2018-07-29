package simulator

import "context"

var CmdRadiusRetry = &CommandSpec{
	[]Token{TRadius, TRetry}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoRadiusRetry = &CommandSpec{
	[]Token{TNo, TRadius, TRetry}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

