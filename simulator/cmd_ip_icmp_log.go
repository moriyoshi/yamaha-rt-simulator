package simulator

import "context"

var CmdIpIcmpLog = &CommandSpec{
	[]Token{TIp, TIcmp, TLog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpIcmpLog = &CommandSpec{
	[]Token{TNo, TIp, TIcmp, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

