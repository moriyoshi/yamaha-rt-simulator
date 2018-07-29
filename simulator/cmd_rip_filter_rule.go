package simulator

import "context"

var CmdRipFilterRule = &CommandSpec{
	[]Token{TRip, TFilter, TRule}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoRipFilterRuleRule = &CommandSpec{
	[]Token{TNo, TRip, TFilter, TRule}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

