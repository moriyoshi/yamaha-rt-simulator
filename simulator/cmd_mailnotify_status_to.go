package simulator

import "context"

var CmdMailNotifyStatusTo = &CommandSpec{
	[]Token{TMailNotify, TStatus, TTo}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMailNotifyStatusToId = &CommandSpec{
	[]Token{TNo, TMailNotify, TStatus, TTo}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

