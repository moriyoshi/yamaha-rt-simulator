package simulator

import "context"

var CmdFrDe = &CommandSpec{
	[]Token{TFr, TDe}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoFrDe = &CommandSpec{
	[]Token{TNo, TFr, TDe}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

