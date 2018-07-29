package simulator

import "context"

var CmdIpv6InterfaceSecureFilter = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TSecure, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpSecureFilter = &CommandSpec{
	[]Token{TIpv6, TPp, TSecure, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6TunnelSecureFilter = &CommandSpec{
	[]Token{TIpv6, TTunnel, TSecure, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceSecureFilter = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TSecure, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpSecureFilter = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TSecure, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6TunnelSecureFilter = &CommandSpec{
	[]Token{TNo, TIpv6, TTunnel, TSecure, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

