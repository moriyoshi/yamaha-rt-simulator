package simulator

import "context"

var CmdSyslogExecuteCommand = &CommandSpec{
	[]Token{TSyslog, TExecute, TCommand}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSyslogExecuteCommand = &CommandSpec{
	[]Token{TNo, TSyslog, TExecute, TCommand}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

