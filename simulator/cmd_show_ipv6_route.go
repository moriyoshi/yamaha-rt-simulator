package simulator

import "context"

var CmdShowIpv6Route = &CommandSpec{
	[]Token{TShow, TIpv6, TRoute}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
