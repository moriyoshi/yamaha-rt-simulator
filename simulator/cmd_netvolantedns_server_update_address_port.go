package simulator

import "context"

var CmdNetvolanteDnsServerUpdateAddressPort = &CommandSpec{
	[]Token{TNetvolanteDns, TServer, TUpdate, TAddress, TPort}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsServerUpdateAddressPort = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, TServer, TUpdate, TAddress, TPort}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

