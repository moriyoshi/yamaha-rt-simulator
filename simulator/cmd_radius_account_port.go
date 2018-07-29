package simulator

import "context"

var CmdRadiusAccountPort = &CommandSpec{
	[]Token{TRadius, TAccount, TPort}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoRadiusAccountPort = &CommandSpec{
	[]Token{TNo, TRadius, TAccount, TPort}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

