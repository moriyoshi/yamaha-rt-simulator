package simulator

import "context"

var CmdIpv6SourceAddressSelectionRule = &CommandSpec{
	[]Token{TIpv6, TSource, TAddress, TSelection, TRule}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6SourceAddressSelectionRule = &CommandSpec{
	[]Token{TNo, TIpv6, TSource, TAddress, TSelection, TRule}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
