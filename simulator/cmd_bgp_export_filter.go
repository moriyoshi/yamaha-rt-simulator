package simulator

import "context"

var CmdBgpExportFilter = &CommandSpec{
	[]Token{TBgp, TExport, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpExportFilter = &CommandSpec{
	[]Token{TNo, TBgp, TExport, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

