package simulator

import "context"

var CmdSipArriveRingingPNUatype = &CommandSpec{
	[]Token{TSip, TArrive, TRinging, TPNUatype}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipArriveRingingPNUatype = &CommandSpec{
	[]Token{TNo, TSip, TArrive, TRinging, TPNUatype}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

