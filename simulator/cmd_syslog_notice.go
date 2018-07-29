package simulator

import "context"

var CmdSyslogNotice = &CommandSpec{
	[]Token{TSyslog, TNotice}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSyslogNotice = &CommandSpec{
	[]Token{TNo, TSyslog, TNotice}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

