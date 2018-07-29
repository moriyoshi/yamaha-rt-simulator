package simulator

import "context"

var CmdSwitchControlConfigSet = &CommandSpec{
	[]Token{TSwitch, TControl, TConfig, TSet, TSomething /* switch */}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
