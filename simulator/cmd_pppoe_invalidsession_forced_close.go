package simulator

import "context"

var CmdPppoeInvalidSessionForcedClose = &CommandSpec{
	[]Token{TPppoe, TInvalidSession, TForced, TClose}, 4,
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

var CmdNoPppoeInvalidSessionForcedClose = &CommandSpec{
	[]Token{TNo, TPppoe, TInvalidSession, TForced, TClose}, 5,
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
