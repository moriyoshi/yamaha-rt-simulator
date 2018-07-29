package simulator

import "context"

var CmdAlarmHttpUpload = &CommandSpec{
	[]Token{TAlarm, THttp, TUpload}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAlarmHttpUpload = &CommandSpec{
	[]Token{TNo, TAlarm, THttp, TUpload}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

