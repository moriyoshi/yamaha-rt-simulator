package simulator

import "context"

var CmdScheduleAt = &CommandSpec{
	[]Token{TSchedule, TAt}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoScheduleAt = &CommandSpec{
	[]Token{TNo, TSchedule, TAt}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
