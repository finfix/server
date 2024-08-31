package repository

import "context"

func (repo *Repository) ChangeSerialNumbers(ctx context.Context, accountGroupID, oldValue, newValue uint32) error {

	var req string
	args := []any{
		accountGroupID,
	}

	if newValue < oldValue {
		req = `UPDATE coin.accounts
			SET serial_number = serial_number + 1
			WHERE account_group_id = ? 
			  AND serial_number >= ?
			  AND serial_number < ?`
		args = append(args, newValue, oldValue)
	} else {
		req = `UPDATE coin.accounts
			SET serial_number = serial_number - 1
			WHERE account_group_id = ? 
			  AND serial_number > ?
			  AND serial_number <= ?`
		args = append(args, oldValue, newValue)
	}

	return repo.db.Exec(ctx, req, args...)
}
