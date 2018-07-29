package simulator

import "context"

var CmdOspfMergeEqualCostStub = &CommandSpec{
	[]Token{TOspf, TMerge, TEqual, TCost, TStub}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOspfMergeEqualCostStub = &CommandSpec{
	[]Token{TNo, TOspf, TMerge, TEqual, TCost, TStub}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

