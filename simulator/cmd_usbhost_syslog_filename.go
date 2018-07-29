package simulator

import "context"

var CmdUsbhostSyslogFilename = &CommandSpec{
	[]Token{TUsbhost, TSyslog, TFilename}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUsbhostSyslogFilename = &CommandSpec{
	[]Token{TNo, TUsbhost, TSyslog, TFilename}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

