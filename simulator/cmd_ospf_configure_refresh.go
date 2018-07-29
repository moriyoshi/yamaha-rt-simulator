package simulator

import "context"

var CmdOspfConfigureRefresh = &CommandSpec{
	[]Token{TOspf, TConfigure, TRefresh}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

