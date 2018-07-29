package simulator

import "context"

var CmdDhcpServerRfc2131Compliant = &CommandSpec{
	[]Token{TDhcp, TServer, TRfc2131, TCompliant}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpServerRfc2131Compliant = &CommandSpec{
	[]Token{TNo, TDhcp, TServer, TRfc2131, TCompliant}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
