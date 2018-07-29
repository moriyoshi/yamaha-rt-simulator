package simulator

import "context"

var CmdIpv6InterfaceRtadvSend = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TRtadv, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpRtadvSend = &CommandSpec{
	[]Token{TIpv6, TPp, TRtadv, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceRtadvSend = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TRtadv, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpRtadvSend = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TRtadv, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

