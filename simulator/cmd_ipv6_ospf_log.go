package simulator

import "context"

var CmdIpv6OspfLog = &CommandSpec{
	[]Token{TIpv6, TOspf, TLog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6OspfLog = &CommandSpec{
	[]Token{TNo, TIpv6, TOspf, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

