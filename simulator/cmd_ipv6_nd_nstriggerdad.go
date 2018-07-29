package simulator

import "context"

var CmdIpv6NdNsTriggerDad = &CommandSpec{
	[]Token{TIpv6, TNd, TNsTriggerDad}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6NdNsTriggerDad = &CommandSpec{
	[]Token{TNo, TIpv6, TNd, TNsTriggerDad}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
