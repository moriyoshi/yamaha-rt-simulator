package simulator

import "context"

var CmdIpv6OspfExportFromOspf = &CommandSpec{
	[]Token{TIpv6, TOspf, TExport, TFrom, TOspf}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6OspfExportFromOspf = &CommandSpec{
	[]Token{TNo, TIpv6, TOspf, TExport, TFrom, TOspf}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

