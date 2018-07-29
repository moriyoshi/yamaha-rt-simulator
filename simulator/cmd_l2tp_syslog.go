package simulator

import "context"

var CmdL2TpSyslog = &CommandSpec{
	[]Token{TL2Tp, TSyslog}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoL2TpSyslog = &CommandSpec{
	[]Token{TNo, TL2Tp, TSyslog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

