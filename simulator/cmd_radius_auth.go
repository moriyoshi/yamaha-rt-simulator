package simulator

import "context"

var CmdRadiusAuth = &CommandSpec{
	[]Token{TRadius, TAuth}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoRadiusAuth = &CommandSpec{
	[]Token{TNo, TRadius, TAuth}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
