package simulator

import "context"

var CmdPppChapMaxchallenge = &CommandSpec{
	[]Token{TPpp, TChap, TMaxchallenge}, 3,
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

var CmdNoPppChapMaxchallenge = &CommandSpec{
	[]Token{TNo, TPpp, TChap, TMaxchallenge}, 4,
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
