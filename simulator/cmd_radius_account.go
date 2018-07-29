package simulator

import "context"

var CmdRadiusAccount = &CommandSpec{
	[]Token{TRadius, TAccount}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoRadiusAccount = &CommandSpec{
	[]Token{TNo, TRadius, TAccount}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

