package simulator

import "context"

var CmdAlarmBatch = &CommandSpec{
	[]Token{TAlarm, TBatch}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAlarmBatch = &CommandSpec{
	[]Token{TNo, TAlarm, TBatch}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

