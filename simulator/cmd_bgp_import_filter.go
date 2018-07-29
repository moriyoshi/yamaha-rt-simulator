package simulator

import "context"

var CmdBgpImportFilter = &CommandSpec{
	[]Token{TBgp, TImport, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpImportFilter = &CommandSpec{
	[]Token{TNo, TBgp, TImport, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

