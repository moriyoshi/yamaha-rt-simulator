package simulator

import "context"

var CmdSyslogLocalAddress = &CommandSpec{
	[]Token{TSyslog, TLocal, TAddress}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSyslogLocalAddress = &CommandSpec{
	[]Token{TNo, TSyslog, TLocal, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

