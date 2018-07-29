package simulator

import "context"

var CmdShowIpv6MrouteFib = &CommandSpec{
	[]Token{TShow, TIpv6, TMroute, TFib}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

