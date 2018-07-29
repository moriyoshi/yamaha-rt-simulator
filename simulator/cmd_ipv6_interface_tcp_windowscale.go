package simulator

import "context"

var CmdIpv6InterfaceTcpWindowScale = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TTcp, TWindowScale}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpTcpWindowScale = &CommandSpec{
	[]Token{TIpv6, TPp, TTcp, TWindowScale}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6TunnelTcpWindowScale = &CommandSpec{
	[]Token{TIpv6, TTunnel, TTcp, TWindowScale}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceTcpWindowScale = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TTcp, TWindowScale}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpTcpWindowScale = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TTcp, TWindowScale}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6TunnelTcpWindowScale = &CommandSpec{
	[]Token{TNo, TIpv6, TTunnel, TTcp, TWindowScale}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

