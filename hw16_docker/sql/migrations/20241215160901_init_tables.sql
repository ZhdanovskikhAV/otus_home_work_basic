-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists Users(
 id uuid primary key default general.new_uuid(),
 name VARCHAR(255) NOT NULL,
 email VARCHAR(255) NOT NULL,
 password VARCHAR(255) NOT NULL
);

CREATE TABLE if not exists Orders (
  id  uuid primary key default general.new_uuid(),
  user_id uuid references Users(id) not null,
  order_date timestamptz not null DEFAULT now(),
  total_amount numeric(10, 2) not null default 0.00
);
CREATE TABLE if not exists Products (
  id  uuid primary key default general.new_uuid(),
  name VARCHAR(255) NOT NULL,
  price numeric(10, 2) not null default 0.00
);

CREATE TABLE if not exists OrderProducts (
    id  uuid primary key default general.new_uuid(),
    product_id uuid references Products(id) not null,
    order_id uuid references Orders(id) not null
);

CREATE INDEX if not exists indx_orders on Orders(user_id, total_amount);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX if exists indx_orders;
DROP TABLE if exists OrderProducts;
DROP TABLE if exists Orders;
DROP TABLE if exists Users;
DROP TABLE if exists Products;

-- +goose StatementEnd