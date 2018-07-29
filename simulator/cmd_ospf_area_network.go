package simulator

import "context"

var CmdOspfAreaNetwork = &CommandSpec{
	[]Token{TOspf, TArea, TNetwork}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOspfAreaNetwork = &CommandSpec{
	[]Token{TNo, TOspf, TArea, TNetwork}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

