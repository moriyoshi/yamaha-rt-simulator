package simulator

import "context"

var CmdShowStatusSwitchControl = &CommandSpec{
	[]Token{TShow, TStatus, TSwitch, TControl}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

