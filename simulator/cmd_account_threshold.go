package simulator

import "context"

var CmdAccountThreshold = &CommandSpec{
	[]Token{TAccount, TThreshold, TPp}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAccountThreshold = &CommandSpec{
	[]Token{TNo, TAccount, TThreshold, TPp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
