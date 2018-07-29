package simulator

import "context"

var CmdApControlConfigGet = &CommandSpec{
	[]Token{TAp, TControl, TConfig, TGet}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
