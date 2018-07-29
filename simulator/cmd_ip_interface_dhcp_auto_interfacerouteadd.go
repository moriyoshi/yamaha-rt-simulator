package simulator

import "context"

var CmdIpInterfaceDhcpAutoInterfaceRouteAdd = &CommandSpec{
	[]Token{TIp, TL2Interface, TDhcp, TAuto, TInterfaceRouteAdd}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceDhcpAutoInterfaceRouteAdd = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TDhcp, TAuto, TInterfaceRouteAdd}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

