package simulator

import "context"

var CmdSwitchControlFunctionDefault = &CommandSpec{
	[]Token{TSwitch, TControl, TFunction, TDefault}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

