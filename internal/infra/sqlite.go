package infra

import (
	"fmt"
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/colfarl/budgy/internal/core"
	"github.com/colfarl/budgy/internal/database"
	"github.com/mattn/go-sqlite3"
)

type SqliteRunner struct{
	DB		*sql.DB
	Q		*database.Queries
}

var ErrNilUsername = errors.New("SQLITE RUNNER: cannot set nil username")

func (sr SqliteRunner) Run(ctx context.Context, fx core.Effect, emit func(core.Event)) bool {
	switch v := fx.(type) {
	case core.FxLoadSession:
		return sr.sqliteLoadSession(ctx, emit)	

	case core.FxSetSessionUser:
		return sr.sqliteSetUser(ctx, v,  emit)	

	case core.FxClearSessionUser:
		return sr.sqliteClearUser(ctx, emit)	

	case core.FxLoadAllUsers:
		return sr.sqliteLoadUsers(ctx, emit)	

	case core.FxCreateUser:
		return sr.sqliteCreateUser(ctx, v, emit)	

	case core.FxDeleteUser:
		return sr.sqliteDeleteUser(ctx, v, emit)	
	
	case core.FxGetUserBalances:
		return sr.sqliteSumUserAccounts(ctx, v, emit)	

	// ============================== Account ==============================
	case core.FxCreateAccount:
		return sr.sqliteCreateAccount(ctx, v, emit)	

	case core.FxDeleteAccount:
		return sr.sqliteDeleteAccount(ctx, v, emit)	

	case core.FxLoadAccounts:
		return sr.sqliteLoadAccounts(ctx, v, emit)

	case core.FxGetAccountBalance:
		return sr.sqliteSumAccount(ctx, v, emit)	
	// ============================== Transactions ==============================
	case core.FxCreateTxn:
		return sr.sqliteCreateTxn(ctx, v, emit)	

	case core.FxDeleteTxn:
		return sr.sqliteDeleteTxn(ctx, v, emit)	

	case core.FxLoadAccountTxns:
		return sr.sqliteLoadAccountTxns(ctx, v, emit)	

	case core.FxImportTxnsFromFile:
		return sr.sqliteImportTxnsFromFile(ctx, v, emit)	

	case core.FxCategorizeTxn:
		return sr.sqliteCategorizeTxn(ctx, v, emit)	

	case core.FxUncategorizeTxn:
		return sr.sqliteUncategorizeTxn(ctx, v, emit)	

	case core.FxSplitTransaction:
		return sr.sqliteSplitTxn(ctx, v, emit)	

	// ============================== Categories ==============================
	case core.FxCreateCategory:
		return sr.sqliteCreateCategory(ctx, v, emit)

	case core.FxDeleteCategory:
		return sr.sqliteDeleteCategory(ctx, v, emit)

	case core.FxLoadCategories:
		return sr.sqliteLoadCategories(ctx, emit)
	}
	return false
}

func (sr SqliteRunner) sqliteClearUser(ctx context.Context, emit func(core.Event)) bool {	
	if err := sr.Q.LogoutSession(ctx); err != nil {
		emit(core.DBFailure{Err: err})
		return true 
	}

	emit(core.ActiveUserCleared{})
	return true
}

func (sr SqliteRunner) sqliteSetUser(ctx context.Context, fx core.FxSetSessionUser, emit func(core.Event)) bool {
	if fx.Username == nil {
		emit(core.DBFailure{Err: ErrNilUsername})
		return true 
	}
	err := sr.Q.LoginSession(ctx, sql.NullString{String: *fx.Username, Valid: true})
	if err != nil {
		emit(core.DBFailure{Err: err})
		return false
	}

	emit(core.ActiveUserSet{Username: *fx.Username})
	return true
}

func (sr SqliteRunner) sqliteLoadUsers(ctx context.Context, emit func(core.Event)) bool {
	users, err := sr.Q.GetAllUserNames(ctx)	
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true 
	}	
	
	emit(core.UsersLoaded{Usernames: users})
	return true
}

func (sr SqliteRunner) sqliteDeleteUser(ctx context.Context, fx core.FxDeleteUser, emit func(core.Event)) bool {
	log.Printf("running effect delete user")
	if fx.Username == nil {
		emit(core.DBFailure{Err: ErrNilUsername})
		return true
	}

	rows, err := sr.Q.DeleteUserByName(ctx, *fx.Username)	
	if err != nil {
		emit(core.DBFailure{Err: ErrNilUsername})
		return true
	} else if rows == 0 {
		// Unsure if this should be handled 
		log.Printf("ATTEMPTED TO DELETE NOT EXISTENT USER")
		return true
	}
	
	emit(core.UserDeleted{Username: *fx.Username})
	return true
}

func (sr SqliteRunner) sqliteCreateUser(ctx context.Context, fx core.FxCreateUser, emit func(core.Event)) bool {
	if fx.Username == nil {
		emit(core.DBFailure{Err: ErrNilUsername})
		return true
	}
	user, err := sr.Q.CreateUser(ctx, *fx.Username)	
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				// unique constraint violation do nothing
				return true 
			}
		} 
		emit(core.DBFailure{Err: ErrNilUsername})
		return true
	}		

	emit(core.UserCreated{Username: user.Name})
	return true
}

func (sr SqliteRunner) sqliteLoadSession(ctx context.Context, emit func(core.Event)) bool {
	prev, err := sr.Q.LoadSession(ctx)	
	if err != nil || !prev.Valid {
		emit(core.SessionLoadFailed{Err: err})
		return true 
	}

	if err = sr.Q.UpdateSessionLastOpened(ctx); err != nil {
		emit(core.DBFailure{Err: err})
		return true 
	}
	
	emit(core.ActiveUserSet{Username: prev.String})
	return true
}

func (sr SqliteRunner) sqliteCreateAccount(ctx context.Context, fx core.FxCreateAccount, emit func(core.Event)) bool {
	user, err := sr.Q.GetUserByName(ctx, fx.Username)	
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true 
	}
	params := database.CreateAccountParams{
		Name: fx.AccountName,	
		UserID: user.ID,
	}

	account, err := sr.Q.CreateAccount(ctx, params)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}

	values := core.AccountCreated{
		UserID: user.ID,	
		Username: user.Name,	
		AccountName: account.Name,	
		AccountID: account.ID,	
	}
	emit(values)
	return true
}

func (sr SqliteRunner) sqliteDeleteAccount(ctx context.Context, fx core.FxDeleteAccount, emit func(core.Event)) bool {
	user, err := sr.Q.GetUserByName(ctx, fx.Username)	
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true 
	}

	params := database.DeleteAccountParams{
		Name: fx.AccountName,	
		UserID: user.ID,
	}

	err = sr.Q.DeleteAccount(ctx, params)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}

	values := core.AccountDeleted{
		UserID: user.ID,	
		Username: user.Name,	
		AccountName: fx.AccountName,	
	}
	emit(values)
	return true
}

func (sr SqliteRunner) sqliteLoadAccounts(ctx context.Context, fx core.FxLoadAccounts, emit func(core.Event)) bool {
	user, err := sr.Q.GetUserByName(ctx, fx.Username)	
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true 
	}

	accounts, err := sr.Q.GetAllAccounts(ctx, user.ID)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}

	values := make([]string, len(accounts))
	for i := range accounts {
		values[i] = accounts[i].Name
	}
	emit(core.AccountsLoaded{AccountNames: values})
	return true
}

func (sr SqliteRunner) sqliteCreateTxn(ctx context.Context, fx core.FxCreateTxn, emit func(core.Event)) bool {	
	params := database.CreateTransactionFromNamesParams{
		Username: fx.Username,
		AccountName: fx.AccountName,
		Amount: fx.Amount,
		Description: fx.Description,
		OccurredAt: fx.Date,
		IsIncome: fx.Income,
	}

	txn, err := sr.Q.CreateTransactionFromNames(ctx, params)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}
	
	emit(core.TxnCreated{
		Transaction: core.Txn{
			ID: txn.ID,
			Username: fx.Username,
			AccountName: fx.AccountName,
			Income: txn.IsIncome,
			Description: txn.Description,
			Amount: txn.Amount,
		},
	})
	return true 
}

func (sr SqliteRunner) sqliteDeleteTxn(ctx context.Context, fx core.FxDeleteTxn, emit func(core.Event)) bool {
	err := sr.Q.DeleteTransaction(ctx, fx.TxnID)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}
	emit(core.TxnDeleted{})
	return true 
}

func (sr SqliteRunner) sqliteLoadAccountTxns(ctx context.Context, fx core.FxLoadAccountTxns, emit func(core.Event)) bool {		
		
	var txns []database.Transaction
	var err error 

	if fx.Uncategorized {
		log.Println("in uncategorized")
		params := database.GetAccountUncategorizedTxnFromNamesParams{
			Username: fx.Username,
			AccountName: fx.AccountName,
		}
		txns, err = sr.Q.GetAccountUncategorizedTxnFromNames(ctx, params)
		log.Println("Got", len(txns))
		if err != nil {
			emit(core.DBFailure{Err: err})
			return true
		}
	} else {
		params := database.GetAccountTxnFromNamesParams{
			Username: fx.Username,
			AccountName: fx.AccountName,
		}
		txns, err = sr.Q.GetAccountTxnFromNames(ctx, params)
		log.Println("Got", len(txns))
		if err != nil {
			emit(core.DBFailure{Err: err})
			return true
		}
	}

	values := make([]core.Txn, len(txns))
	for i := range txns {
		values[i] = core.Txn{
			ID: txns[i].ID,
			Username: fx.Username,
			AccountName: fx.AccountName,
			Amount: txns[i].Amount,
			Description: txns[i].Description,
			Income: txns[i].IsIncome,
		}
	}

	emit(core.AccountTxnsLoaded{Transactions: values})
	return true 
}

func (sr SqliteRunner) sqliteImportTxnsFromFile(ctx context.Context, fx core.FxImportTxnsFromFile, emit func(core.Event)) bool {
	key := FileKey{Bank: fx.FileOrigin, Format: fx.FileType}
	reader, ok := FileRegistry[key]	
	if !ok {
		emit(core.GeneralFailure{Err: ErrUnsupportedFile})
		return true
	}
	
	txns, err := reader(fx.FilePath)
	if err != nil {
		emit(core.GeneralFailure{Err: err})
		return true
	}
	
	uploaded := make([]core.Txn, 0)
	for _, v := range txns {
		params := database.CreateTransactionFromNamesParams{
			Username: fx.Username,
			AccountName: fx.AccountName,
			IsIncome: v.Income,
			Amount: v.Amount,
			Description: v.Description,
			OccurredAt: v.Date,
		}

		inserted, err := sr.Q.CreateTransactionFromNames(ctx, params)
		if err != nil {
			// fail silently
			log.Printf("skipping: %v\n", err)
			continue
		}

		uploaded = append(uploaded, core.Txn{
			ID: inserted.ID,
			Username: fx.Username,
			AccountName: fx.AccountName,
			Amount: inserted.Amount, 
			Income: inserted.IsIncome, 
		})
	}

	emit(core.TxnsImported{Transactions: uploaded})
	return true
}

func (sr SqliteRunner) sqliteCreateCategory(ctx context.Context, fx core.FxCreateCategory, emit func(core.Event)) bool {
	params := database.CreateCategoryParams{
		Name: fx.Category.Name,
		IsIncome: fx.Category.IsIncome,
	}

	cat, err := sr.Q.CreateCategory(ctx, params)
	if err != nil {
		emit(core.GeneralFailure{Err: err})
		return true
	}
	
	emit(core.CategoryCreated{Category: core.Category{
		ID: cat.ID,
		Name: cat.Name, 
		IsIncome: cat.IsIncome,
	}})
	return true
}

func (sr SqliteRunner) sqliteDeleteCategory(ctx context.Context, fx core.FxDeleteCategory, emit func(core.Event)) bool {
	err := sr.Q.DeleteCategory(ctx, fx.ID)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}

	emit(core.CategoryDeleted{ID: fx.ID})
	return true
}

func (sr SqliteRunner) sqliteLoadCategories(ctx context.Context, emit func(core.Event)) bool {
	cats, err := sr.Q.GetAllCategories(ctx) 
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}
	
	values := make([]core.Category, len(cats))
	for i := range cats {
		values[i] = core.Category{
			ID: cats[i].ID,
			Name: cats[i].Name,
			IsIncome: cats[i].IsIncome,
		}
	}

	emit(core.CategoriesLoaded{Categories: values})
	return true
}

func (sr SqliteRunner) sqliteUncategorizeTxn(ctx context.Context, fx core.FxUncategorizeTxn, emit func(core.Event)) bool {	
	err := sr.Q.UncategorizeTransaction(ctx, fx.ID)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}

	emit(core.TxnUncategorized{ID: fx.ID})
	return true
}

func (sr SqliteRunner) sqliteCategorizeTxn(ctx context.Context, fx core.FxCategorizeTxn, emit func(core.Event)) bool {	
	params := database.CategorizeTransactionByNameParams{
		TransactionID: fx.ID,
		Name: fx.Category,
	}

	err := sr.Q.CategorizeTransactionByName(ctx, params)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}

	emit(core.TxnCategorized{ID: fx.ID, Category: fx.Category})
	return true
}

func (sr SqliteRunner) sqliteSplitTxn(ctx context.Context, fx core.FxSplitTransaction, emit func(core.Event)) bool {	

	original, err := sr.Q.GetTxnFromID(ctx, fx.ID)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}
	
	category_lookup := make(map[string]struct{})
	categories, err := sr.Q.GetAllCategories(ctx)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}
	
	for _, c := range categories {
		category_lookup[c.Name] = struct{}{}
	}
	
	tx, err := sr.DB.Begin()
	defer tx.Rollback()

	total := 0.0

	params := database.CreateTransactionParams{
		AccountID: original.AccountID,
		IsIncome: original.IsIncome,
		Amount: original.Amount,
		Description: original.Description,	
		OccurredAt: original.OccurredAt,
	}
	
	result := make([][]string, 0)
	for _, s := range fx.Splits {
		new_cat := ""
		amount := s 		

		if i := strings.IndexByte(s, ':'); i >= 0 {
			amount = s[:i]
			new_cat = s[i + 1:] 
		}

		amount_money, err := strconv.ParseFloat(amount, 64)	
		if err != nil {	
			emit(core.GeneralFailure{Err: err})
			return true
		}
		total += amount_money 
		params.Amount = amount_money

		txn, err := sr.Q.CreateTransaction(ctx, params)
		if err != nil {	
			emit(core.DBFailure{Err: err})
			return true
		}
		
		new_cat = strings.ToLower(new_cat)
		if new_cat != "" {
			if _, ok := category_lookup[new_cat]; !ok {
				emit(core.GeneralFailure{
					Err: fmt.Errorf("Category %v does not exist", new_cat),
				})
				return true
			}

			categorize_params := database.CategorizeTransactionByNameParams{
				TransactionID: txn.ID,
				Name: new_cat,
			}

			err = sr.Q.CategorizeTransactionByName(ctx, categorize_params)
			if err != nil {	
				emit(core.DBFailure{Err: err})
				return true
			}
		}
		result = append(result, []string{amount, new_cat})
	}

	if total != original.Amount {
		emit(core.GeneralFailure{
			Err: fmt.Errorf("New amount != original amount"),
		})
		return true
	}
	
	err = sr.Q.DeleteTransaction(ctx, original.ID)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}
	
	core_original := core.Txn{
		ID: original.ID,
		Amount: original.Amount,
		Description: original.Description,
		Income: original.IsIncome,  
	}
	emit(core.TxnSplit{Old: core_original, New: result})
	tx.Commit()
	return true
}

func (sr SqliteRunner) sqliteSumAccount(ctx context.Context, fx core.FxGetAccountBalance, emit func(core.Event)) bool {	
	params := database.SumAccountTxnFromNamesParams{
		Username: fx.Username,
		AccountName: fx.AccountName,
	}

	sum, err := sr.Q.SumAccountTxnFromNames(ctx, params)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}

	if !sum.Valid {
		emit(core.GeneralFailure{Err: err})
		return true
	}	
	
	emit(core.AccountSummed{AccountName: fx.AccountName, AccountSum: sum.Float64})
	return true
}

func (sr SqliteRunner) sqliteSumUserAccounts(ctx context.Context, fx core.FxGetUserBalances, emit func(core.Event)) bool {	
	sums, err := sr.Q.SumAccountFromUsername(ctx, fx.Username)
	if err != nil {
		emit(core.DBFailure{Err: err})
		return true
	}
	
	result := make([]core.SummedAccount, len(sums))
	for i, v := range sums {
		if !v.Balance.Valid {
			emit(core.DBFailure{Err: err})
			return true
		}

		result[i] = core.SummedAccount{
			Name: v.Name,
			Balance: v.Balance.Float64,
		}
	}
			
	emit(core.UserSummed{Accounts: result})
	return true
}

