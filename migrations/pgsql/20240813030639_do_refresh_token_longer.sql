-- +goose Up
-- +goose StatementBegin
ALTER TABLE coin.devices ALTER COLUMN refresh_token TYPE varchar(400) USING refresh_token::varchar(400);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE coin.devices ALTER COLUMN refresh_token TYPE varchar(250) USING refresh_token::varchar(250);
-- +goose StatementEnd
