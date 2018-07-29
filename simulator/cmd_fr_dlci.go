package simulator

import "context"

var CmdFrDlci = &CommandSpec{
	[]Token{TFr, TDlci}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoFrDlci = &CommandSpec{
	[]Token{TNo, TFr, TDlci}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

