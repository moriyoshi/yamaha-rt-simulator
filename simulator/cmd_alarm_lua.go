package simulator

import "context"

var CmdAlarmLua = &CommandSpec{
	[]Token{TAlarm, TLua}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoAlarmLua = &CommandSpec{
	[]Token{TNo, TAlarm, TLua}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

