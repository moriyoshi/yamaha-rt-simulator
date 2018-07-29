package simulator

import "context"

var CmdBgpImport = &CommandSpec{
	[]Token{TBgp, TImport}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpImport = &CommandSpec{
	[]Token{TNo, TBgp, TImport}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

