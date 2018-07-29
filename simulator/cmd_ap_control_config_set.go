package simulator

import "context"

var CmdApControlConfigSet = &CommandSpec{
	[]Token{TAp, TControl, TConfig, TSet}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
