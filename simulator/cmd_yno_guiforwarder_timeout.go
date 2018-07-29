package simulator

import "context"

var CmdYnoGuiForwarderTimeout = &CommandSpec{
	[]Token{TYno, TGuiForwarder, TTimeout}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoYnoGuiForwarderTimeout = &CommandSpec{
	[]Token{TNo, TYno, TGuiForwarder, TTimeout}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

