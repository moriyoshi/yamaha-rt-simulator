package simulator

import "context"

var CmdIpv6InterfaceOspfArea = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TOspf, TArea}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpOspfArea = &CommandSpec{
	[]Token{TIpv6, TPp, TOspf, TArea}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6TunnelOspfArea = &CommandSpec{
	[]Token{TIpv6, TTunnel, TOspf, TArea}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceOspfArea = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TOspf, TArea}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpOspfArea = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TOspf, TArea}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6TunnelOspfArea = &CommandSpec{
	[]Token{TNo, TIpv6, TTunnel, TOspf, TArea}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

