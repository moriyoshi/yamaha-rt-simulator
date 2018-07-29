package simulator

import "context"

var CmdIpPolicyFilterSetSwitch = &CommandSpec{
	[]Token{TIp, TPolicy, TFilter, TSet, TSwitch}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
