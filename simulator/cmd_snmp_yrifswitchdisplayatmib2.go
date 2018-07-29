package simulator

import "context"

var CmdSnmpYrifswitchdisplayatmib2 = &CommandSpec{
	[]Token{TSnmp, TYrifswitchdisplayatmib2}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpYrifswitchdisplayatmib2 = &CommandSpec{
	[]Token{TNo, TSnmp, TYrifswitchdisplayatmib2}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

