package simulator

import "context"

var CmdOspfRouterId = &CommandSpec{
	[]Token{TOspf, TRouter, TId}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOspfRouterId = &CommandSpec{
	[]Token{TNo, TOspf, TRouter, TId}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

