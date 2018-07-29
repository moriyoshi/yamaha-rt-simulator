package simulator

import "context"

var CmdMobileArriveUse = &CommandSpec{
	[]Token{TMobile, TArrive, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileArriveUse = &CommandSpec{
	[]Token{TNo, TMobile, TArrive, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

