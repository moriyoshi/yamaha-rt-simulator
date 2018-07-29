package simulator

import "context"

var CmdYnoHttpsProxy = &CommandSpec{
	[]Token{TYno, THttpsProxy}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoYnoHttpsProxy = &CommandSpec{
	[]Token{TNo, TYno, THttpsProxy}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

