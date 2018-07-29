package simulator

import "context"

var CmdIpInterfaceTcpMssLimit = &CommandSpec{
	[]Token{TIp, TL2Interface, TTcp, TMss, TLimit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpTcpMssLimit = &CommandSpec{
	[]Token{TIp, TPp, TTcp, TMss, TLimit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelTcpMssLimit = &CommandSpec{
	[]Token{TIp, TTunnel, TTcp, TMss, TLimit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceTcpMssLimit = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TTcp, TMss, TLimit}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpTcpMssLimit = &CommandSpec{
	[]Token{TNo, TIp, TPp, TTcp, TMss, TLimit}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelTcpMssLimit = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TTcp, TMss, TLimit}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

