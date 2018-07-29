package simulator

import "context"

var CmdMailNotifyStatusTimeout = &CommandSpec{
	[]Token{TMailNotify, TStatus, TTimeout}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMailNotifyStatusTimeout = &CommandSpec{
	[]Token{TNo, TMailNotify, TStatus, TTimeout}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

