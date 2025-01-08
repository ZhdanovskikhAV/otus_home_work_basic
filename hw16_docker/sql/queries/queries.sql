-- name: CreateUser :one
insert into  users (name, email, password)
VALUES ($1, $2, $3)
returning id;

-- name: CreateProduct :one
insert into products (name, price)
VALUES ($1, $2)
returning id;

-- name: UpdateUser :one
update users
SET name = $2,
    email = $3
WHERE id = $1
returning id;

-- name: UpdateProduct :one
update products
SET name = $2,
    price = $3
WHERE id = $1
returning id;

-- name: DeleteUser :one
delete from users
WHERE id = $1
returning id;

-- name: DeleteProduct :one
delete from products
WHERE id = $1
returning id;

-- name: CreateOrder :one
insert into  orders (user_id, order_date, total_amount)
VALUES ($1, $2, $3)
returning id;

-- name: CreateOrderWithCurrentDate :one
insert into  orders (user_id, total_amount)
VALUES ($1, $2)
returning id;

-- name: DeleteOrder :one
delete from orders
WHERE id = $1
returning id;

-- name: GetUserById :one
select * from users where id = $1;

-- name: GetProductById :one
select * from products where id = $1;

-- name: GetProductByName :one
select * from products where name = $1;

-- name: SortProductByPriceASC :many
select * from products order by price asc limit $1;

-- name: SortProductByPriceDESC :many
select * from products order by price desc limit $1;

-- name: SordUserById :many
select * from users order by id limit $1;

-- name: GetOrdersByUser :many
select * from orders where user_id = $1;

-- name: GetUserTotalPrice :one
select sum(total_amount) from orders where user_id = $1;

-- name: GetUserAvaragePrice :one
select avg(total_amount) from orders where user_id = $1;

-- name: GetAllUsers :many
select * from users;

-- name: GetAllProducts :many
select * from products;

-- name: GetAllOrders :many
select * from orders;

-- name: GetUserByName :one
select * from users where name = $1;