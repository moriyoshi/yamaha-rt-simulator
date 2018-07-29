package simulator

import "context"

var CmdEthernetInterfaceFilter = &CommandSpec{
	[]Token{TEthernet, TL2Interface, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoEthernetInterfaceFilter = &CommandSpec{
	[]Token{TNo, TEthernet, TL2Interface, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

