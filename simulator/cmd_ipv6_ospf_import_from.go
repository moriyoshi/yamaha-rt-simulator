package simulator

import "context"

var CmdIpv6OspfImportFrom = &CommandSpec{
	[]Token{TIpv6, TOspf, TImport, TFrom}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6OspfImportFrom = &CommandSpec{
	[]Token{TNo, TIpv6, TOspf, TImport, TFrom}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

