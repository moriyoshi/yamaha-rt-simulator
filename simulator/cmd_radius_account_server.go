package simulator

import "context"

var CmdRadiusAccountServer = &CommandSpec{
	[]Token{TRadius, TAccount, TServer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoRadiusAccountServer = &CommandSpec{
	[]Token{TNo, TRadius, TAccount, TServer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
