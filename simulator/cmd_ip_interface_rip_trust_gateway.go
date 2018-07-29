package simulator

import "context"

var CmdIpInterfaceRipTrustGateway = &CommandSpec{
	[]Token{TIp, TL2Interface, TRip, TTrust, TGateway}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpRipTrustGateway = &CommandSpec{
	[]Token{TIp, TPp, TRip, TTrust, TGateway}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelRipTrustGateway = &CommandSpec{
	[]Token{TIp, TTunnel, TRip, TTrust, TGateway}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceRipTrustGateway = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TRip, TTrust, TGateway}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpRipTrustGateway = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, TTrust, TGateway}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelRipTrustGateway = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TRip, TTrust, TGateway}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

