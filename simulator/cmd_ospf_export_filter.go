package simulator

import "context"

var CmdOspfExportFilter = &CommandSpec{
	[]Token{TOspf, TExport, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOspfExportFilter = &CommandSpec{
	[]Token{TNo, TOspf, TExport, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

