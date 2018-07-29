package simulator

import "context"

var CmdIpPolicyFilterTimer = &CommandSpec{
	[]Token{TIp, TPolicy, TFilter, TTimer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPolicyFilterTimer = &CommandSpec{
	[]Token{TNo, TIp, TPolicy, TFilter, TTimer}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

