package simulator

import "context"

var CmdMobileSyslog = &CommandSpec{
	[]Token{TMobile, TSyslog}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileSyslog = &CommandSpec{
	[]Token{TNo, TMobile, TSyslog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

