package simulator

import "context"

var CmdMobileCarrierMode = &CommandSpec{
	[]Token{TMobile, TCarrier, TMode}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileCarrierMode = &CommandSpec{
	[]Token{TNo, TMobile, TCarrier, TMode}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

