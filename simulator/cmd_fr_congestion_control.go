package simulator

import "context"

var CmdFrCongestionControl = &CommandSpec{
	[]Token{TFr, TCongestion, TControl}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoFrCongestionControl = &CommandSpec{
	[]Token{TNo, TFr, TCongestion, TControl}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

