package simulator

import "context"

var CmdSnmpTrapMobileSignalStrength = &CommandSpec{
	[]Token{TSnmp, TTrap, TMobile, TSignalStrength}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpTrapMobileSignalStrength = &CommandSpec{
	[]Token{TNo, TSnmp, TTrap, TMobile, TSignalStrength}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

