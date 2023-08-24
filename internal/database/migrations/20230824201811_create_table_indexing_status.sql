-- +goose Up
-- +goose StatementBegin
create table indexing_status(
    id uuid primary key default gen_random_uuid(),
    entity varchar unique not null,
    last_block int default -1 not null
);
insert into indexing_status (entity) values ('orders');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table indexing_status;
-- +goose StatementEnd
