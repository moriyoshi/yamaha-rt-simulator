package simulator

import "context"

var CmdNgnRadiusAccountCaller = &CommandSpec{
	[]Token{TNgn, TRadius, TAccount, TCaller}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNgnRadiusAccountCaller = &CommandSpec{
	[]Token{TNo, TNgn, TRadius, TAccount, TCaller}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

