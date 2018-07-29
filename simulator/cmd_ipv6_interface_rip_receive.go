package simulator

import "context"

var CmdIpv6InterfaceRipReceive = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TRip, TReceive}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpRipReceive = &CommandSpec{
	[]Token{TIpv6, TPp, TRip, TReceive}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6TunnelRipReceive = &CommandSpec{
	[]Token{TIpv6, TTunnel, TRip, TReceive}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceRipReceive = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TRip, TReceive}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpRipReceive = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TRip, TReceive}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6TunnelRipReceive = &CommandSpec{
	[]Token{TNo, TIpv6, TTunnel, TRip, TReceive}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

