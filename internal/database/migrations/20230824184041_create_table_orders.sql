-- +goose Up
-- +goose StatementBegin
create table orders (
    id uuid primary key default gen_random_uuid(),
    user_address character(42) not null,
    value numeric not null,
    tx_hash character(66) unique not null,
    block_time timestamp not null,
    created_at timestamp default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table orders;
-- +goose StatementEnd
