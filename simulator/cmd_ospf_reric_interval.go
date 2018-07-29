package simulator

import "context"

var CmdOspfRericInterval = &CommandSpec{
	[]Token{TOspf, TReric, TInterval}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOspfRericInterval = &CommandSpec{
	[]Token{TNo, TOspf, TReric, TInterval}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

