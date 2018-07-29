package simulator

import "context"

var CmdPptpCallIdMode = &CommandSpec{
	[]Token{TPptp, TCallId, TMode}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoPptpCallIdMode = &CommandSpec{
	[]Token{TNo, TPptp, TCallId, TMode}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
