package simulator

import "context"

var CmdQueueClassFilter = &CommandSpec{
	[]Token{TQueue, TClass, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdQueueClassFilterPrecedence = &CommandSpec{
	[]TokensArityPair{
		{[]Token{TQueue, TClass, TFilter, TSomething /* num */, TPrecedence, TMappingArg, TCosArg, &EitherToken{TIp, TIpv6}}, 8},
		{[]Token{TQueue, TClass, TFilter, TSomething /* num */, TPrecedence, TCosArg, &EitherToken{TIp, TIpv6}}, 7},
		{[]Token{TQueue, TClass, TFilter, TSomething /* num */, TPrecedence, TMappingArg, &EitherToken{TIp, TIpv6}}, 7},
		{[]Token{TQueue, TClass, TFilter, TSomething /* num */, TPrecedence, &EitherToken{TIp, TIpv6}}, 6},
	}, 0,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdQueueClassFilterDscp = &CommandSpec{
	[]TokensArityPair{
		{[]Token{TQueue, TClass, TFilter, TSomething /* num */, TDscp, TCosArg, &EitherToken{TIp, TIpv6}}, 5},
		{[]Token{TQueue, TClass, TFilter, TSomething /* num */, TDscp, &EitherToken{TIp, TIpv6}}, 5},
	}, 0,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueueClassFilter = &CommandSpec{
	[]Token{TNo, TQueue, TClass, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
