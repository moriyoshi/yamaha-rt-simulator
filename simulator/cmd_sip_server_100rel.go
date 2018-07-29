package simulator

import "context"

var CmdSipServer100Rel = &CommandSpec{
	[]Token{TSip, TServer, T100Rel}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipServer100Rel = &CommandSpec{
	[]Token{TNo, TSip, TServer, T100Rel}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

