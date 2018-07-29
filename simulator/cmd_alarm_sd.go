package simulator

import "context"

var CmdAlarmSd = &CommandSpec{
	[]Token{TAlarm, TSd}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAlarmSdSwitch = &CommandSpec{
	[]Token{TNo, TAlarm, TSd}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

