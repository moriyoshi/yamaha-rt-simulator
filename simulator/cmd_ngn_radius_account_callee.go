package simulator

import "context"

var CmdNgnRadiusAccountCallee = &CommandSpec{
	[]Token{TNgn, TRadius, TAccount, TCallee}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNgnRadiusAccountCallee = &CommandSpec{
	[]Token{TNo, TNgn, TRadius, TAccount, TCallee}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

