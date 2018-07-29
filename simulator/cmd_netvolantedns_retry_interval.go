package simulator

import "context"

var CmdNetvolanteDnsRetryInterval = &CommandSpec{
	[]Token{TNetvolanteDns, TRetry, TInterval}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNetvolanteDnsRetryIntervalPp = &CommandSpec{
	[]Token{TNetvolanteDns, TRetry, TInterval, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsRetryInterval = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, TRetry, TInterval}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsRetryIntervalPpIntervalCount = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, TRetry, TInterval, TPp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

