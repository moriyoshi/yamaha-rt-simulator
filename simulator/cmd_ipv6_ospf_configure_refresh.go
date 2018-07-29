package simulator

import "context"

var CmdIpv6OspfConfigureRefresh = &CommandSpec{
	[]Token{TIpv6, TOspf, TConfigure, TRefresh}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

