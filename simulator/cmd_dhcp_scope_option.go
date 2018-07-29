package simulator

import "context"

var CmdDhcpScopeOption = &CommandSpec{
	[]Token{TDhcp, TScope, TOption}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpScopeOption = &CommandSpec{
	[]Token{TNo, TDhcp, TScope, TOption}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

