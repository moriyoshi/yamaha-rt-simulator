package simulator

import "context"

var CmdIpv6OspfPreference = &CommandSpec{
	[]Token{TIpv6, TOspf, TPreference}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6OspfPreference = &CommandSpec{
	[]Token{TNo, TIpv6, TOspf, TPreference}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

