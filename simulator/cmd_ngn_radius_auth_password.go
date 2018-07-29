package simulator

import "context"

var CmdNgnRadiusAuthPassword = &CommandSpec{
	[]Token{TNgn, TRadius, TAuth, TPassword}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNgnRadiusAuthPassword = &CommandSpec{
	[]Token{TNo, TNgn, TRadius, TAuth, TPassword}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

