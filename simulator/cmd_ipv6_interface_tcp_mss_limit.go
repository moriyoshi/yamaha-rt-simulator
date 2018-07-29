package simulator

import "context"

var CmdIpv6InterfaceTcpMssLimit = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TTcp, TMss, TLimit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpTcpMssLimit = &CommandSpec{
	[]Token{TIpv6, TPp, TTcp, TMss, TLimit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6TunnelTcpMssLimit = &CommandSpec{
	[]Token{TIpv6, TTunnel, TTcp, TMss, TLimit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceTcpMssLimit = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TTcp, TMss, TLimit}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpTcpMssLimit = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TTcp, TMss, TLimit}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6TunnelTcpMssLimit = &CommandSpec{
	[]Token{TNo, TIpv6, TTunnel, TTcp, TMss, TLimit}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

