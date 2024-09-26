package repository

import "context"

// ChangeSerialNumbers вставляет группу счетов на новое место
func (repo *AccountGroupRepository) ChangeSerialNumbers(ctx context.Context, oldValue, newValue uint32) error {

	var req string
	var args []any

	if newValue < oldValue {
		req = `UPDATE coin.account_groups
			SET serial_number = serial_number + 1
			WHERE serial_number >= ?
			  AND serial_number < ?`
		args = append(args, newValue, oldValue)
	} else {
		req = `UPDATE coin.account_groups
			SET serial_number = serial_number - 1
			WHERE serial_number > ?
			  AND serial_number <= ?`
		args = append(args, oldValue, newValue)
	}

	return repo.db.Exec(ctx, req, args...)
}
