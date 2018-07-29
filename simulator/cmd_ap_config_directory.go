package simulator

import "context"

var CmdApConfigDirectory = &CommandSpec{
	[]Token{TAp, TConfig, TDirectory}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoApConfigDirectory = &CommandSpec{
	[]Token{TNo, TAp, TConfig, TDirectory}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

