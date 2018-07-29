package simulator

import "context"

var CmdShowStatusEthernetFilter = &CommandSpec{
	[]Token{TShow, TStatus, TEthernet, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

