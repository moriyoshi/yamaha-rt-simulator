package simulator

import "context"

var CmdShowIpv6RipTable = &CommandSpec{
	[]Token{TShow, TIpv6, TRip, TTable}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

