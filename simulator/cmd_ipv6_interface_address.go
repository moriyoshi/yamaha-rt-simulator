package simulator

import "context"

var CmdIpv6InterfaceAddress = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpAddress = &CommandSpec{
	[]Token{TIpv6, TPp, TAddress}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6TunnelAddress = &CommandSpec{
	[]Token{TIpv6, TTunnel, TAddress}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
var CmdNoIpv6InterfaceAddress = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpAddress = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6TunnelAddress = &CommandSpec{
	[]Token{TNo, TIpv6, TTunnel, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
