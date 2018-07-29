package simulator

import "context"

var CmdShowStatusIpv6Mld = &CommandSpec{
	[]Token{TShow, TStatus, TIpv6, TMld}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

