package simulator

import "context"

var CmdIpInterfaceTrafficListThreshold = &CommandSpec{
	[]Token{TIp, TL2Interface, TTraffic, TList, TThreshold}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpTrafficListThreshold = &CommandSpec{
	[]Token{TIp, TPp, TTraffic, TList, TThreshold}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelTrafficListThreshold = &CommandSpec{
	[]Token{TIp, TTunnel, TTraffic, TList, TThreshold}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceTrafficListThreshold = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TTraffic, TList, TThreshold}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpTrafficListThreshold = &CommandSpec{
	[]Token{TNo, TIp, TPp, TTraffic, TList, TThreshold}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelTrafficListThreshold = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TTraffic, TList, TThreshold}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

