package simulator

var messages = map[string]map[interface{}]string{
	"ja": {
		"Error: %s\n":                  "エラー：%s\n",
		InvalidCommandName:             "コマンド名を確認してください",
		IllegalKeyword:                 "パラメータのキーワードが認識できません",
		AdministratorUseOnly:           "このコマンドは管理レベルでのみ使用できます",
		WrongNumberOfArgs:              "パラメータの数が不適当です",
		IncorrectPassword:              "パスワードが間違っています",
		UnsavedChangesInVolatileMemory: "不揮発性メモリに保存されていない設定変更があります",
		PromptSaveSettings:             "新しい設定を保存しますか? (Y/N)",
		PPSelectionRequired:            "このコマンドは相手を選択してから実行してください",
		IntegerArgRequired:             "パラメータは整数で指定してください",
		ArgOutOfRange:                  "パラメータが範囲を越えています",
		UnrecognizedIPAddr:             "IPアドレスが認識できません",
	},
	"en": {
		"Error: %s\n":                  "Error: %s\n",
		InvalidCommandName:             "Invalid command name",
		IllegalKeyword:                 "Illegal keyword",
		AdministratorUseOnly:           "Administrator use only",
		WrongNumberOfArgs:              "Insufficient or too many parameters",
		IncorrectPassword:              "Incorrect password",
		UnsavedChangesInVolatileMemory: "Unsaved changes in volatile memory",
		PromptSaveSettings:             "Save new configuration ? (Y/N)",
		PPSelectionRequired:            "Execute pp select command beforehand",
		IntegerArgRequired:             "Parameter must be an integer",
		ArgOutOfRange:                  "Parameter out of range",
		UnrecognizedIPAddr:             "Unrecognized IP address parameter",
	},
}
