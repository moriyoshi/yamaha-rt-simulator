package simulator

import "context"

var CmdShowNatDescriptorAddress = &CommandSpec{
	[]Token{TShow, TNat, TDescriptor, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

