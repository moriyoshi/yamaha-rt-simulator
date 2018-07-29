package simulator

import "context"

var CmdLanLinkAggregationStatic = &CommandSpec{
	[]Token{TLan, TLinkAggregation, TStatic}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanLinkAggregationStatic = &CommandSpec{
	[]Token{TNo, TLan, TLinkAggregation, TStatic}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

