package simulator

import "context"

var CmdAlarmEntire = &CommandSpec{
	[]Token{TAlarm, TEntire}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAlarmEntire = &CommandSpec{
	[]Token{TNo, TAlarm, TEntire}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

