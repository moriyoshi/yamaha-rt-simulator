package simulator

import "context"

var CmdSyslogInfo = &CommandSpec{
	[]Token{TSyslog, TInfo}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSyslogInfo = &CommandSpec{
	[]Token{TNo, TSyslog, TInfo}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

