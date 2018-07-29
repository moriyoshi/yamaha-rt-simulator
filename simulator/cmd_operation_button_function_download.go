package simulator

import "context"

var CmdOperationButtonFunctionDownload = &CommandSpec{
	[]Token{TOperation, TButton, TFunction, TDownload}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOperationButtonFunctionDownload = &CommandSpec{
	[]Token{TNo, TOperation, TButton, TFunction, TDownload}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

