package simulator

import "context"

var CmdVlanPortMapping = &CommandSpec{
	[]Token{TVlan, TPort, TMapping}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoVlanPortMapping = &CommandSpec{
	[]Token{TNo, TVlan, TPort, TMapping}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

