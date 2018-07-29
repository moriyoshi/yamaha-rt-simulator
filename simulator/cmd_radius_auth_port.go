package simulator

import "context"

var CmdRadiusAuthPort = &CommandSpec{
	[]Token{TRadius, TAuth, TPort}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoRadiusAuthPort = &CommandSpec{
	[]Token{TNo, TRadius, TAuth, TPort}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

