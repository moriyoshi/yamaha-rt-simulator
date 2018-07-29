package simulator

import "context"

var CmdMailNotifyStatusFrom = &CommandSpec{
	[]Token{TMailNotify, TStatus, TFrom}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMailNotifyStatusFrom = &CommandSpec{
	[]Token{TNo, TMailNotify, TStatus, TFrom}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

