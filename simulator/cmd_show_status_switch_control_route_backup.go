package simulator

import "context"

var CmdShowStatusSwitchControlRouteBackup = &CommandSpec{
	[]Token{TShow, TStatus, TSwitch, TControl, TRoute, TBackup}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

