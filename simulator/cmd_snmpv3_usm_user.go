package simulator

import "context"

var CmdSnmpv3UsmUser = &CommandSpec{
	[]Token{TSnmpv3, TUsm, TUser, TSomething /* user_id */, TSomething /* name */, TGroup}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpv3UsmUser = &CommandSpec{
	[]Token{TNo, TSnmpv3, TUsm, TUser}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

