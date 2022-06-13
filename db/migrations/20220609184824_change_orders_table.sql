-- +goose Up
-- +goose StatementBegin
ALTER TABLE statuses RENAME COLUMN status_id TO id;

ALTER TABLE orders_history RENAME COLUMN history_id TO id;

ALTER TABLE orders_history RENAME COLUMN id_order TO order_id;

ALTER TABLE cab_mans RENAME COLUMN cab_man_id TO id;

ALTER TABLE orders RENAME COLUMN id_status TO status_id;

ALTER TABLE orders RENAME COLUMN id_cab_man TO cab_man_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
