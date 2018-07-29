package simulator

import "context"

var CmdAlarmUsbhost = &CommandSpec{
	[]Token{TAlarm, TUsbhost}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAlarmUsbhost = &CommandSpec{
	[]Token{TNo, TAlarm, TUsbhost}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

