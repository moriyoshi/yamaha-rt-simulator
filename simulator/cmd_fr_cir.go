package simulator

import "context"

var CmdFrCir = &CommandSpec{
	[]Token{TFr, TCir}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoFrCir = &CommandSpec{
	[]Token{TNo, TFr, TCir}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

