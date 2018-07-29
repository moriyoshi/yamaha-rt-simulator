package simulator

import "context"

var CmdBgpExport = &CommandSpec{
	[]Token{TBgp, TExport, TAspath}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpExport = &CommandSpec{
	[]Token{TNo, TBgp, TExport, TAspath}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
