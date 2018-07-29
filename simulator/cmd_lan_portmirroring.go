package simulator

import "context"

var CmdLanPortMirroring = &CommandSpec{
	[]Token{TLan, TPortMirroring}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanPortMirroring = &CommandSpec{
	[]Token{TNo, TLan, TPortMirroring}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

