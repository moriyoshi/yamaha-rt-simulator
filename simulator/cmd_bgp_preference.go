package simulator

import "context"

var CmdBgpPreference = &CommandSpec{
	[]Token{TBgp, TPreference}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpPreference = &CommandSpec{
	[]Token{TNo, TBgp, TPreference}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

