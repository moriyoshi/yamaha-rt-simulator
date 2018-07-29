package simulator

import "context"

var CmdSwitchControlConfigGet = &CommandSpec{
	[]Token{TSwitch, TControl, TConfig, TGet}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdSwitchControlConfigGetInterface = &CommandSpec{
	[]Token{TSwitch, TControl, TConfig, TGet}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

