package simulator

import "context"

var CmdAlarmMobile = &CommandSpec{
	[]Token{TAlarm, TMobile}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAlarmMobile = &CommandSpec{
	[]Token{TNo, TAlarm, TMobile}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

