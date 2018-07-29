package simulator

import "context"

var CmdMailNotifyStatusServer = &CommandSpec{
	[]Token{TMailNotify, TStatus, TServer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMailNotifyStatusServer = &CommandSpec{
	[]Token{TNo, TMailNotify, TStatus, TServer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

