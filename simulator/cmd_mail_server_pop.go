package simulator

import "context"

var CmdMailServerPop = &CommandSpec{
	[]Token{TMail, TServer, TPop}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMailServerPop = &CommandSpec{
	[]Token{TNo, TMail, TServer, TPop}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

