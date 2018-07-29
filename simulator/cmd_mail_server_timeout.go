package simulator

import "context"

var CmdMailServerTimeout = &CommandSpec{
	[]Token{TMail, TServer, TTimeout}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMailServerTimeout = &CommandSpec{
	[]Token{TNo, TMail, TServer, TTimeout}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

