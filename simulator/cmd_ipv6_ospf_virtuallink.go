package simulator

import "context"

var CmdIpv6OspfVirtualLink = &CommandSpec{
	[]Token{TIpv6, TOspf, TVirtualLink}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6OspfVirtualLink = &CommandSpec{
	[]Token{TNo, TIpv6, TOspf, TVirtualLink}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

