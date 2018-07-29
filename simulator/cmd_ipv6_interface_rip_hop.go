package simulator

import "context"

var CmdIpv6InterfaceRipHop = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TRip, THop}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpRipHopDirectionHop = &CommandSpec{
	[]Token{TIpv6, TPp, TRip, THop}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceRipHop = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TRip, THop}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpRipHop = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TRip, THop}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

