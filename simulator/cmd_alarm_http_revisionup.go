package simulator

import "context"

var CmdAlarmHttpRevisionUp = &CommandSpec{
	[]Token{TAlarm, THttp, TRevisionUp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAlarmHttpRevisionUpSwitch = &CommandSpec{
	[]Token{TNo, TAlarm, THttp, TRevisionUp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

