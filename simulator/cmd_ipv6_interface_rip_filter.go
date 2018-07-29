package simulator

import "context"

var CmdIpv6InterfaceRipFilter = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TRip, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpRipFilter = &CommandSpec{
	[]Token{TIpv6, TPp, TRip, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6TunnelRipFilter = &CommandSpec{
	[]Token{TIpv6, TTunnel, TRip, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceRipFilter = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TRip, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpRipFilter = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TRip, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6TunnelRipFilter = &CommandSpec{
	[]Token{TNo, TIpv6, TTunnel, TRip, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

