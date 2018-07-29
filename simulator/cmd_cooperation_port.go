package simulator

import "context"

var CmdCooperationPort = &CommandSpec{
	[]Token{TCooperation, TPort}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoCooperationPort = &CommandSpec{
	[]Token{TNo, TCooperation, TPort}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

