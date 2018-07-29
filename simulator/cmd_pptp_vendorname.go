package simulator

import "context"

var CmdPptpVendorname = &CommandSpec{
	[]Token{TPptp, TVendorname}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoPptpVendorname = &CommandSpec{
	[]Token{TNo, TPptp, TVendorname}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
