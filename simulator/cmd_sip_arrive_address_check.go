package simulator

import "context"

var CmdSipArriveAddressCheck = &CommandSpec{
	[]Token{TSip, TArrive, TAddress, TCheck}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipArriveAddressCheck = &CommandSpec{
	[]Token{TNo, TSip, TArrive, TAddress, TCheck}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

