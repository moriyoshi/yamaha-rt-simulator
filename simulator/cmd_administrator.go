package simulator

import "context"

var CmdAdministrator = &CommandSpec{
	[]Token{TAdministrator}, 1,
	func(_ context.Context, sess *SimulatorSession, _ []TokenInstance) error {
		line, err := sess.Terminal.ReadPassword("Password: ")
		if err != nil {
			return err
		}
		if line != "" {
			return IncorrectPassword
		}
		sess.Enabled = true
		return nil
	},
}
