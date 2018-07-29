package simulator

import "context"

var CmdDhcpScopeBind = &CommandSpec{
	[]Token{TDhcp, TScope, TBind}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpScopeBind = &CommandSpec{
	[]Token{TNo, TDhcp, TScope, TBind}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
