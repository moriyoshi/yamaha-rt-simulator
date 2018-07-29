package simulator

import "context"

var CmdIsdnArrivePermit = &CommandSpec{
	[]Token{TIsdn, TArrive, TPermit}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnArrivePermit = &CommandSpec{
	[]Token{TNo, TIsdn, TArrive, TPermit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

