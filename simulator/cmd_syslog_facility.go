package simulator

import "context"

var CmdSyslogFacility = &CommandSpec{
	[]Token{TSyslog, TFacility}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSyslogFacility = &CommandSpec{
	[]Token{TNo, TSyslog, TFacility}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

