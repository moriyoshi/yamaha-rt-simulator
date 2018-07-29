package simulator

import "context"

var CmdSnmpYrifppdisplayatmib2 = &CommandSpec{
	[]Token{TSnmp, TYrifppdisplayatmib2}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpYrifppdisplayatmib2 = &CommandSpec{
	[]Token{TNo, TSnmp, TYrifppdisplayatmib2}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

