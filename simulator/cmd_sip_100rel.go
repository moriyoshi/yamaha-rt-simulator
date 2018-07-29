package simulator

import "context"

var CmdSip100Rel = &CommandSpec{
	[]Token{TSip, T100Rel}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSip100Rel = &CommandSpec{
	[]Token{TNo, TSip, T100Rel}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

