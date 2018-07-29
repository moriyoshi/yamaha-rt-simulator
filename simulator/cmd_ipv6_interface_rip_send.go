package simulator

import "context"

var CmdIpv6InterfaceRipSend = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TRip, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpRipSend = &CommandSpec{
	[]Token{TIpv6, TPp, TRip, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6TunnelRipSend = &CommandSpec{
	[]Token{TIpv6, TTunnel, TRip, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceRipSend = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TRip, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpRipSend = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TRip, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6TunnelRipSend = &CommandSpec{
	[]Token{TNo, TIpv6, TTunnel, TRip, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

