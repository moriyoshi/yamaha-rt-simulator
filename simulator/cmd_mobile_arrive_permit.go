package simulator

import "context"

var CmdMobileArrivePermit = &CommandSpec{
	[]Token{TMobile, TArrive, TPermit}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileArrivePermit = &CommandSpec{
	[]Token{TNo, TMobile, TArrive, TPermit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

