-- +goose Up
-- +goose StatementBegin

-- Удаляем лишнюю таблицу
DROP TABLE coin.account_permissions;

-- Добавляем новое поле со связкой с группой счетов для транзакций
ALTER TABLE coin.transactions ADD account_group_id int NULL;
ALTER TABLE coin.transactions ADD CONSTRAINT transactions_account_groups_fk FOREIGN KEY (account_group_id) REFERENCES coin.account_groups(id);

-- Присваиваем каждой транзакции группу счетов по счету отправителя
UPDATE coin.transactions t SET account_group_id = a.account_group_id FROM coin.accounts a WHERE t.account_from_id = a.id;

-- Делаем поле account_group_id обязательным
ALTER TABLE coin.transactions ALTER COLUMN account_group_id SET NOT NULL;

-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin

-- Восстанавливаем таблицу account_permissions
CREATE TABLE coin.account_permissions ();

-- Удаляем поле со связкой с группой счетов для транзакций
ALTER TABLE coin.transactions DROP CONSTRAINT transactions_account_groups_fk;
ALTER TABLE coin.transactions DROP COLUMN account_group_id;

-- +goose StatementEnd
