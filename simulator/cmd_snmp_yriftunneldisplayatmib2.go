package simulator

import "context"

var CmdSnmpYriftunneldisplayatmib2 = &CommandSpec{
	[]Token{TSnmp, TYriftunneldisplayatmib2}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpYriftunneldisplayatmib2 = &CommandSpec{
	[]Token{TNo, TSnmp, TYriftunneldisplayatmib2}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

