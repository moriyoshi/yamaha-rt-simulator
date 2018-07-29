package simulator

import "context"

var CmdIpInterfaceTcpWindowScale = &CommandSpec{
	[]Token{TIp, TL2Interface, TTcp, TWindowScale}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpTcpWindowScale = &CommandSpec{
	[]Token{TIp, TPp, TTcp, TWindowScale}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelTcpWindowScale = &CommandSpec{
	[]Token{TIp, TTunnel, TTcp, TWindowScale}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceTcpWindowScale = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TTcp, TWindowScale}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpTcpWindowScale = &CommandSpec{
	[]Token{TNo, TIp, TPp, TTcp, TWindowScale}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelTcpWindowScale = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TTcp, TWindowScale}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

