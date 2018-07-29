package simulator

import "context"

var CmdIpv6InterfaceRipTrustGateway = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TRip, TTrust, TGateway}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpRipTrustGateway = &CommandSpec{
	[]Token{TIpv6, TPp, TRip, TTrust, TGateway}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceRipTrustGateway = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TRip, TTrust, TGateway}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpRipTrustGateway = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TRip, TTrust, TGateway}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

