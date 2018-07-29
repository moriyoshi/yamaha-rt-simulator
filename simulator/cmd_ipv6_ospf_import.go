package simulator

import "context"

var CmdIpv6OspfImport = &CommandSpec{
	[]Token{TIpv6, TOspf, TImport}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6OspfImport = &CommandSpec{
	[]Token{TNo, TIpv6, TOspf, TImport}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

