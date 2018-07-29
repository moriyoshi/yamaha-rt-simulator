package simulator

import "context"

var CmdIpv6InterfaceMtu = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TMtu}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpMtu = &CommandSpec{
	[]Token{TIpv6, TPp, TMtu}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceMtu = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TMtu}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpMtu = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TMtu}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

