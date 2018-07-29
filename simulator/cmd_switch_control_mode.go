package simulator

import "context"

var CmdSwitchControlMode = &CommandSpec{
	[]Token{TSwitch, TControl, TMode}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSwitchControlMode = &CommandSpec{
	[]Token{TNo, TSwitch, TControl, TMode}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

