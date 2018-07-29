package simulator

import "context"

var CmdEthernetFilter = &CommandSpec{
	[]Token{TEthernet, TFilter}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
