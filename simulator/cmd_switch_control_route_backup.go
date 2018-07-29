package simulator

import "context"

var CmdSwitchControlRouteBackup = &CommandSpec{
	[]Token{TSwitch, TControl, TRoute, TBackup}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSwitchControlRouteBackupRoute = &CommandSpec{
	[]Token{TNo, TSwitch, TControl, TRoute, TBackup}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

