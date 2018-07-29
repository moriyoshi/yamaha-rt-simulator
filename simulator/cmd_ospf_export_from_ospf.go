package simulator

import "context"

var CmdOspfExportFromOspf = &CommandSpec{
	[]Token{TOspf, TExport, TFrom, TOspf}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOspfExportFromOspf = &CommandSpec{
	[]Token{TNo, TOspf, TExport, TFrom, TOspf}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

