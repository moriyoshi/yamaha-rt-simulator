package simulator

import "context"

var CmdSyslogSrcport = &CommandSpec{
	[]Token{TSyslog, TSrcport}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSyslogSrcport = &CommandSpec{
	[]Token{TNo, TSyslog, TSrcport}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

