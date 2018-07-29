package simulator

import "context"

var CmdDnsPrivateAddressSpoof = &CommandSpec{
	[]Token{TDns, TPrivate, TAddress, TSpoof}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDnsPrivateAddressSpoof = &CommandSpec{
	[]Token{TNo, TDns, TPrivate, TAddress, TSpoof}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

