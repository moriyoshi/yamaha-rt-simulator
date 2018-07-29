package simulator

import "context"

var CmdBgpAutonomousSystem = &CommandSpec{
	[]Token{TBgp, TAutonomousSystem}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpAutonomousSystem = &CommandSpec{
	[]Token{TNo, TBgp, TAutonomousSystem}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

