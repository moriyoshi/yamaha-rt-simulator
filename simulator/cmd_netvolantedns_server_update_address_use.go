package simulator

import "context"

var CmdNetvolanteDnsServerUpdateAddressUse = &CommandSpec{
	[]Token{TNetvolanteDns, TServer, TUpdate, TAddress, TUse}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsServerUpdateAddressUse = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, TServer, TUpdate, TAddress, TUse}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

