package simulator

import "context"

var CmdIpsecLogIllegalSpi = &CommandSpec{
	[]Token{TIpsec, TLog, TIllegalSpi}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecLogIllegalSpi = &CommandSpec{
	[]Token{TNo, TIpsec, TLog, TIllegalSpi}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

