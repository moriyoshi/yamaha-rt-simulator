package simulator

import "context"

var CmdSwitchConfigFilename = &CommandSpec{
	[]Token{TSwitch, TConfig, TFilename}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSwitchConfigFilename = &CommandSpec{
	[]Token{TNo, TSwitch, TConfig, TFilename}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

