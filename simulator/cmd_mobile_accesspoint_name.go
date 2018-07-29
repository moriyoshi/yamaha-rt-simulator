package simulator

import "context"

var CmdMobileAccessPointName = &CommandSpec{
	[]Token{TMobile, TAccessPoint, TName}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileAccessPointName = &CommandSpec{
	[]Token{TNo, TMobile, TAccessPoint, TName}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

