package simulator

import "context"

var CmdApControlConfigDelete = &CommandSpec{
	[]Token{TAp, TControl, TConfig, TDelete}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

