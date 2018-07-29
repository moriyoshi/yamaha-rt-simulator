package simulator

import "context"

var CmdIpInterfaceDhcpAutoDefaultRouteAdd = &CommandSpec{
	[]Token{TIp, TL2Interface, TDhcp, TAuto, TDefaultRouteAdd}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceDhcpAutoDefaultRouteAdd = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TDhcp, TAuto, TDefaultRouteAdd}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

