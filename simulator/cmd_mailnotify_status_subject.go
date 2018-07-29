package simulator

import "context"

var CmdMailNotifyStatusSubject = &CommandSpec{
	[]Token{TMailNotify, TStatus, TSubject}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMailNotifyStatusSubject = &CommandSpec{
	[]Token{TNo, TMailNotify, TStatus, TSubject}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

