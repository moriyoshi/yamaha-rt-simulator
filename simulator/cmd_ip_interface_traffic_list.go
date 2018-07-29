package simulator

import "context"

var CmdIpInterfaceTrafficList = &CommandSpec{
	[]Token{TIp, TL2Interface, TTraffic, TList}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpTrafficList = &CommandSpec{
	[]Token{TIp, TPp, TTraffic, TList}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelTrafficList = &CommandSpec{
	[]Token{TIp, TTunnel, TTraffic, TList}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceTrafficList = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TTraffic, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpTrafficList = &CommandSpec{
	[]Token{TNo, TIp, TPp, TTraffic, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelTrafficList = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TTraffic, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

