package simulator

import "context"

var CmdPpAuthMultiConnectProhibit = &CommandSpec{
	[]Token{TPp, TAuth, TMulti, TConnect, TProhibit}, 6,
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

var CmdNoPpAuthMultiConnectProhibit = &CommandSpec{
	[]Token{TNo, TPp, TAuth, TMulti, TConnect, TProhibit}, 6,
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
