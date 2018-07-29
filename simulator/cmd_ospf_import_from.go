package simulator

import "context"

var CmdOspfImportFrom = &CommandSpec{
	[]Token{TOspf, TImport, TFrom}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOspfImportFrom = &CommandSpec{
	[]Token{TNo, TOspf, TImport, TFrom}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

