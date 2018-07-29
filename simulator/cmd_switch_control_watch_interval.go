package simulator

import "context"

var CmdSwitchControlWatchInterval = &CommandSpec{
	[]Token{TSwitch, TControl, TWatch, TInterval}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSwitchControlWatchInterval = &CommandSpec{
	[]Token{TNo, TSwitch, TControl, TWatch, TInterval}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

