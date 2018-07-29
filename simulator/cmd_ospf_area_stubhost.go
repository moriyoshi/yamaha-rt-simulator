package simulator

import "context"

var CmdOspfAreaStubhost = &CommandSpec{
	[]Token{TOspf, TArea, TStubhost}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOspfAreaStubhost = &CommandSpec{
	[]Token{TNo, TOspf, TArea, TStubhost}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

