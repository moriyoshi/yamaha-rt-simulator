package simulator

import "context"

var CmdOspfImportFilter = &CommandSpec{
	[]Token{TOspf, TImport, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOspfImportFilter = &CommandSpec{
	[]Token{TNo, TOspf, TImport, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

