package simulator

import "context"

var CmdLanCountHubOverflow = &CommandSpec{
	[]Token{TLan, TCountHubOverflow}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanCountHubOverflow = &CommandSpec{
	[]Token{TNo, TLan, TCountHubOverflow}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

