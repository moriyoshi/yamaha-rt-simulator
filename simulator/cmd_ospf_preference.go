package simulator

import "context"

var CmdOspfPreference = &CommandSpec{
	[]Token{TOspf, TPreference}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOspfPreference = &CommandSpec{
	[]Token{TNo, TOspf, TPreference}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

