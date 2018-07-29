package simulator

import "context"

var CmdFrLmi = &CommandSpec{
	[]Token{TFr, TLmi}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoFrLmi = &CommandSpec{
	[]Token{TNo, TFr, TLmi}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

