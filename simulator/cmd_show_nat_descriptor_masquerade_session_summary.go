package simulator

import "context"

var CmdShowNatDescriptorMasqueradeSessionSummary = &CommandSpec{
	[]Token{TShow, TNat, TDescriptor, TMasquerade, TSession, TSomething /* nat_descriptor */, TSummary}, 7,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

