package simulator

import "context"

var CmdMailNotifyStatusUse = &CommandSpec{
	[]Token{TMailNotify, TStatus, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMailNotifyStatusUse = &CommandSpec{
	[]Token{TNo, TMailNotify, TStatus, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

