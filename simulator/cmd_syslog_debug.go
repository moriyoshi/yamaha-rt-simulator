package simulator

import "context"

var CmdSyslogDebug = &CommandSpec{
	[]Token{TSyslog, TDebug}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSyslogDebug = &CommandSpec{
	[]Token{TNo, TSyslog, TDebug}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

