package simulator

import "context"

var CmdIpv6InterfaceDhcpService = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TDhcp, TService}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpDhcpService = &CommandSpec{
	[]Token{TIpv6, TPp, TDhcp, TService}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6TunnelDhcpService = &CommandSpec{
	[]Token{TIpv6, TTunnel, TDhcp, TService}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceDhcpService = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TDhcp, TService}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpDhcpService = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TDhcp, TService}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6TunnelDhcpService = &CommandSpec{
	[]Token{TNo, TIpv6, TTunnel, TDhcp, TService}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
