package simulator

import "context"

var CmdIpInterfaceDhcpService = &CommandSpec{
	[]Token{TIp, TL2Interface, TDhcp, TService}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceDhcpService = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TDhcp, TService}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

