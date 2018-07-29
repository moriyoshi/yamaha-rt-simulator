package simulator

import "context"

var CmdIpv6OspfAreaNetwork = &CommandSpec{
	[]Token{TIpv6, TOspf, TArea, TNetwork}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6OspfAreaNetwork = &CommandSpec{
	[]Token{TNo, TIpv6, TOspf, TArea, TNetwork}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

