type CategoryAddCmd struct {
	Name 	string 	`arg:"" help:"User who owns Account."`
}

type CategoryDeleteCmd struct {
	TxnID	int64  	`arg:"" help:"ID of transaction to delete"`
}

type CategoryListCmd struct {
	Username string 	`arg:"" help:"Username of transaction owner"`
	AccountName string 	`arg:"" help:"Account of transaction."`
}
