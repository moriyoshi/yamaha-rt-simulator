package simulator

import "context"

var CmdBgpRericInterval = &CommandSpec{
	[]Token{TBgp, TReric, TInterval}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpReric = &CommandSpec{
	[]Token{TNo, TBgp, TReric, TInterval}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

