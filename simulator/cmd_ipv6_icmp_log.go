package simulator

import "context"

var CmdIpv6IcmpLog = &CommandSpec{
	[]Token{TIpv6, TIcmp, TLog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6IcmpLog = &CommandSpec{
	[]Token{TNo, TIpv6, TIcmp, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

