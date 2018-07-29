package simulator

import "context"

var CmdIsdnCallBlockTime = &CommandSpec{
	[]Token{TIsdn, TCall, TBlock, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnCallBlockTime = &CommandSpec{
	[]Token{TNo, TIsdn, TCall, TBlock, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

