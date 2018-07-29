package simulator

import "context"

var CmdApConfigFilename = &CommandSpec{
	[]Token{TAp, TConfig, TFilename}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoApConfigFilename = &CommandSpec{
	[]Token{TNo, TAp, TConfig, TFilename}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

