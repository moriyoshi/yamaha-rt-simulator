package simulator

import "context"

var CmdPppPapMaxauthreq = &CommandSpec{
	[]Token{TPpp, TPap, TMaxauthreq}, 3,
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

var CmdNoPppPapMaxauthreq = &CommandSpec{
	[]Token{TNo, TPpp, TPap, TMaxauthreq}, 4,
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
