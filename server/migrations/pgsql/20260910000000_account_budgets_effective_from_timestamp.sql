-- +goose Up
-- +goose StatementBegin
-- effective_from был DATE — Postgres отбрасывал время из присланного клиентом Date.now, поэтому
-- несколько версий бюджета, созданных в один день, после синка получали ОДИНАКОВЫЙ effective_from
-- (00:00:00) и порядок между ними становился неопределённым: резолв "действующей версии" мог
-- взять не последнюю. Клиент и так шлёт/принимает полный timestamp (google.protobuf.Timestamp),
-- обрезала только колонка. Делаем её TIMESTAMPTZ — теперь каждая версия дня уникальна по времени.
ALTER TABLE coin.account_budgets
    ALTER COLUMN effective_from TYPE TIMESTAMP WITH TIME ZONE
    USING effective_from::timestamptz;
COMMENT ON COLUMN coin.account_budgets.effective_from IS 'Дата и время, с которых действует эта версия бюджета';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE coin.account_budgets
    ALTER COLUMN effective_from TYPE DATE
    USING effective_from::date;
COMMENT ON COLUMN coin.account_budgets.effective_from IS 'Дата, с которой действует эта версия бюджета';
-- +goose StatementEnd
