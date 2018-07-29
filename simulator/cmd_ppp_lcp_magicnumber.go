package simulator

import "context"

var CmdPppLcpMagicnumber = &CommandSpec{
	[]Token{TPpp, TLcp, TMagicnumber}, 3,
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

var CmdNoPppLcpMagicnumber = &CommandSpec{
	[]Token{TNo, TPpp, TLcp, TMagicnumber}, 4,
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
