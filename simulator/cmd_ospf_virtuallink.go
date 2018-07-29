package simulator

import "context"

var CmdOspfVirtualLink = &CommandSpec{
	[]Token{TOspf, TVirtualLink}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOspfVirtualLink = &CommandSpec{
	[]Token{TNo, TOspf, TVirtualLink}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

