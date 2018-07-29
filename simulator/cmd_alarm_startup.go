package simulator

import "context"

var CmdAlarmStartup = &CommandSpec{
	[]Token{TAlarm, TStartup}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAlarmStartup = &CommandSpec{
	[]Token{TNo, TAlarm, TStartup}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

