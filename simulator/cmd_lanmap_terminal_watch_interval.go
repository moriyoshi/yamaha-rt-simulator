package simulator

import "context"

var CmdLanMapTerminalWatchInterval = &CommandSpec{
	[]Token{TLanMap, TTerminal, TWatch, TInterval}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanMapTerminalWatchInterval = &CommandSpec{
	[]Token{TNo, TLanMap, TTerminal, TWatch, TInterval}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

