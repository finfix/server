-- +goose Up
-- +goose StatementBegin
CREATE TABLE coin.transaction_files
(
    id                 INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    transaction_id     INT                              NULL,
    original_file_name VARCHAR                          NOT NULL,
    file_name          VARCHAR                          NOT NULL,
    created_by_user_id INT                              NOT NULL,
    datetime_create    timestamptz                      NOT NULL,
    CONSTRAINT transaction_files_unique UNIQUE (id),
    CONSTRAINT transaction_files_transactions_fk FOREIGN KEY (transaction_id) REFERENCES coin.transactions (id),
    CONSTRAINT transaction_files_users_fk FOREIGN KEY (created_by_user_id) REFERENCES coin.users (id)
);

-- Column comments

COMMENT ON COLUMN coin.transaction_files.id IS 'Идентификатор файла';
COMMENT ON COLUMN coin.transaction_files.transaction_id IS 'Идентификатор привязанной транзакции';
COMMENT ON COLUMN coin.transaction_files.original_file_name IS 'Оригинальное название файла';
COMMENT ON COLUMN coin.transaction_files.file_name IS 'Уникальное название файла в виде UUID';
COMMENT ON COLUMN coin.transaction_files.created_by_user_id IS 'Какой пользователь создал файл';
COMMENT ON COLUMN coin.transaction_files.datetime_create IS 'Дата и время создания файла';
COMMENT ON TABLE coin.transaction_files IS 'Файлы транзакций';
COMMENT ON TABLE coin.users_to_account_groups IS 'Доступы пользователей к группам счетов';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE coin.transaction_files;
-- +goose StatementEnd
