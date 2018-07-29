package simulator

import "context"

var CmdFrInarp = &CommandSpec{
	[]Token{TFr, TInarp}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoFrInarp = &CommandSpec{
	[]Token{TNo, TFr, TInarp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

