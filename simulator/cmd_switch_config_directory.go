package simulator

import "context"

var CmdSwitchConfigDirectory = &CommandSpec{
	[]Token{TSwitch, TConfig, TDirectory}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSwitchConfigDirectoryPath = &CommandSpec{
	[]Token{TNo, TSwitch, TConfig, TDirectory}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

