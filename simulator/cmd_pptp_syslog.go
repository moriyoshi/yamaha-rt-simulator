package simulator

import "context"

var CmdPptpSyslog = &CommandSpec{
	[]Token{TPptp, TSyslog}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoPptpSyslog = &CommandSpec{
	[]Token{TNo, TPptp, TSyslog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
