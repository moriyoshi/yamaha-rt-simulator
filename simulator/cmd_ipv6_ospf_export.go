package simulator

import "context"

var CmdIpv6OspfExport = &CommandSpec{
	[]Token{TIpv6, TOspf, TExport}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6OspfExport = &CommandSpec{
	[]Token{TNo, TIpv6, TOspf, TExport}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

