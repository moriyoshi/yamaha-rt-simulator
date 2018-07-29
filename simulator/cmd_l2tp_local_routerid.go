package simulator

import "context"

var CmdL2TpLocalRouterId = &CommandSpec{
	[]Token{TL2Tp, TLocal, TRouterId}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoL2TpLocalRouterId = &CommandSpec{
	[]Token{TNo, TL2Tp, TLocal, TRouterId}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

