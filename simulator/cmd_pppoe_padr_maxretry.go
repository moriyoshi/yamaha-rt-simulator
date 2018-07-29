package simulator

import "context"

var CmdPppoePadrMaxretry = &CommandSpec{
	[]Token{TPppoe, TPadr, TMaxretry}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}

var CmdNoPppoePadrMaxretry = &CommandSpec{
	[]Token{TNo, TPppoe, TPadr, TMaxretry}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}
