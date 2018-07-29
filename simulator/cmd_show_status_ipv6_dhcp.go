package simulator

import "context"

var CmdShowStatusIpv6Dhcp = &CommandSpec{
	[]Token{TShow, TStatus, TIpv6, TDhcp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

