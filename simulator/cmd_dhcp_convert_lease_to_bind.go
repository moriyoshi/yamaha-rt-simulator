package simulator

import "context"

var CmdDhcpConvertLeaseToBind = &CommandSpec{
	[]Token{TDhcp, TConvert, TLease, TTo, TBind}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

