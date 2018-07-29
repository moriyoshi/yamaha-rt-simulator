package simulator

import "context"

var CmdDhcpScopeLeaseType = &CommandSpec{
	[]Token{TDhcp, TScope, TLease, TType}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpScopeLeaseType = &CommandSpec{
	[]Token{TNo, TDhcp, TScope, TLease, TType}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
