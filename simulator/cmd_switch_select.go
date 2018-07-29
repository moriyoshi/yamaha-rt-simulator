package simulator

import "context"

var CmdSwitchSelect = &CommandSpec{
	[]Token{TSwitch, TSelect}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSwitchSelect = &CommandSpec{
	[]Token{TNo, TSwitch, TSelect}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

