package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw15_go_sql/internal/repository/product"
	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw15_go_sql/pkg/pgdb"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// GetUser  получает пользователя по имени из базы данных.
func GetUser(ctx context.Context, dsn string, name string) (product.User, error) {
	// Устанавливаем соединение с базой данных.
	dbc, err := pgdb.New(ctx, dsn, 1)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return product.User{}, err
	}
	defer dbc.Close() // Закрываем соединение после выполнения функции.

	// Создаем экземпляр requestor для работы с пользователями.
	requestor := product.New(dbc)
	user, err := requestor.GetUserByName(ctx, name) // Получаем пользователя по имени.
	if err != nil {
		fmt.Printf("failed to get user by name: %v", err)
		return product.User{}, err
	}
	return *user, nil // Возвращаем найденного пользователя.
}

// CreateUser  создает нового пользователя в базе данных.
func CreateUser(ctx context.Context, dsn string, name string, email string, password string) (string, error) {
	// Устанавливаем соединение с базой данных.
	dbc, err := pgdb.New(ctx, dsn, 1)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return "", err
	}
	defer dbc.Close() // Закрываем соединение после выполнения функции.

	// Начинаем транзакцию.
	tx, err := dbc.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		fmt.Printf("failed to start transaction: %v", err)
		return "", err
	}
	defer func() {
		err = tx.Rollback(ctx) // Откатываем транзакцию в случае ошибки.
		if err != nil {
			fmt.Printf("failed to rollback transaction: %v", err)
		}
	}()

	// Создаем экземпляр requestor для работы с пользователями.
	requestor := product.New(tx)
	requestor.WithTx(tx) // Устанавливаем транзакцию в requestor.

	// Получаем всех пользователей для проверки существования.
	users, err := requestor.GetAllUsers(ctx)
	if err != nil {
		fmt.Printf("failed to get all users: %v", err)
		return "", err
	}

	// Проверяем, существует ли пользователь с таким именем.
	for _, user := range users {
		if user.Name == name {
			errorString := fmt.Sprintf("user with name %v already exists with id: %v", user.Name, user.ID)
			fmt.Println(errorString)
			err = errors.New(errorString)
			return "", err
		}
	}

	// Создаем нового пользователя.
	userID, err := requestor.CreateUser(ctx, product.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		fmt.Printf("failed to create user: %v", err)
		return "", err
	}

	// Коммитим транзакцию.
	err = tx.Commit(ctx)
	if err != nil {
		fmt.Printf("failed to commit transaction: %v", err)
		return "", err
	}
	return userID.String(), nil // Возвращаем ID нового пользователя.
}

// GetOrdersByUser  получает заказы пользователя по имени.
func GetOrdersByUser(ctx context.Context, dsn string, userName string) ([]*product.Order, error) {
	var userID uuid.UUID
	// Устанавливаем соединение с базой данных.
	dbc, err := pgdb.New(ctx, dsn, 1)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return nil, err
	}
	defer dbc.Close() // Закрываем соединение после выполнения функции.

	// Начинаем транзакцию.
	tx, err := dbc.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		fmt.Printf("failed to start transaction: %v", err)
		return []*product.Order{}, err
	}
	defer func() {
		err = tx.Rollback(ctx) // Откатываем транзакцию в случае ошибки.
		if err != nil {
			fmt.Printf("failed to rollback transaction: %v", err)
		}
	}()

	// Создаем экземпляр requestor для работы с пользователями.
	requestor := product.New(tx)
	requestor.WithTx(tx)

	// Получаем всех пользователей для поиска пользователя по имени.
	users, err := requestor.GetAllUsers(ctx)
	if err != nil {
		fmt.Printf("failed to get information about users: %v", err)
		return []*product.Order{}, err
	}

	// Ищем пользователя по имени.
	for _, user := range users {
		if user.Name == userName {
			userID = user.ID
			break
		}
	}
	if userID == [16]byte{} { // Проверяем, найден ли пользователь.
		errorString := fmt.Sprintf("user with name %v not found", userName)
		fmt.Println(errorString)
		err = errors.New(errorString)
		return []*product.Order{}, err
	}

	// Получаем заказы пользователя.
	orders, err := requestor.GetOrdersByUser(ctx, userID)
	if err != nil {
		fmt.Printf("failed to get information about orders: %v", err)
		return []*product.Order{}, err
	}

	// Коммитим транзакцию.
	err = tx.Commit(ctx)
	if err != nil {
		fmt.Printf("failed to commit transaction: %v", err)
		return []*product.Order{}, err
	}
	return orders, nil // Возвращаем заказы пользователя.
}

// CreateOrder создает новый заказ для пользователя.
func CreateOrder(ctx context.Context, dsn string, userName string, totalAmount string) error {
	var numeric pgtype.Numeric
	numeric.Scan(totalAmount) // Преобразуем строку в числовой тип.
	// Устанавливаем соединение с базой данных.
	dbc, err := pgdb.New(ctx, dsn, 1)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return err
	}
	defer dbc.Close() // Закрываем соединение после выполнения функции.

	// Начинаем транзакцию.
	tx, err := dbc.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		fmt.Printf("failed to start transaction: %v", err)
		return err
	}
	defer func() {
		err = tx.Rollback(ctx) // Откатываем транзакцию в случае ошибки.
		if err != nil {
			fmt.Printf("failed to rollback transaction: %v", err)
		}
	}()

	// Создаем экземпляр requestor для работы с пользователями.
	requestor := product.New(tx)
	requestor.WithTx(tx)

	// Получаем всех пользователей для поиска пользователя по имени.
	users, err := requestor.GetAllUsers(ctx)
	if err != nil {
		fmt.Printf("failed to get all users: %v", err)
		return err
	}

	// Ищем пользователя по имени.
	for _, user := range users {
		if user.Name == userName {
			// Создаем новый заказ с текущей датой.
			_, err = requestor.CreateOrderWithCurrentDate(ctx, product.CreateOrderWithCurrentDateParams{
				UserID:      user.ID,
				TotalAmount: numeric,
			})
			if err != nil {
				fmt.Printf("failed to create order: %v", err)
				return err
			}
			// Коммитим транзакцию.
			err = tx.Commit(ctx)
			if err != nil {
				fmt.Printf("failed to commit transaction: %v", err)
				return err
			}
			return nil // Успешное создание заказа.
		}
	}
	errorString := fmt.Sprintf("user with name %v not found", userName)
	fmt.Println(errorString)
	err = errors.New(errorString)
	return err // Возвращаем ошибку, если пользователь не найден.
}

// GetProductByName получает продукт по имени.
func GetProductByName(ctx context.Context, dsn string, name string) (product.Product, error) {
	// Устанавливаем соединение с базой данных.
	dbc, err := pgdb.New(ctx, dsn, 1)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return product.Product{}, err
	}
	defer dbc.Close() // Закрываем соединение после выполнения функции.

	// Создаем экземпляр requestor для работы с продуктами.
	requestor := product.New(dbc)
	product, err := requestor.GetProductByName(ctx, name) // Получаем продукт по имени.
	if err != nil {
		fmt.Printf("failed to get product by name: %v", err)
		return *product, err
	}
	return *product, nil // Возвращаем найденный продукт.
}

// CreateProduct создает новый продукт в базе данных.
func CreateProduct(ctx context.Context, dsn string, name string, price string) (string, error) {
	var numeric pgtype.Numeric
	numeric.Scan(price) // Преобразуем строку в числовой тип.
	// Устанавливаем соединение с базой данных.
	dbc, err := pgdb.New(ctx, dsn, 1)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return "", err
	}
	defer dbc.Close() // Закрываем соединение после выполнения функции.

	// Начинаем транзакцию.
	tx, err := dbc.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		fmt.Printf("failed to start transaction: %v", err)
		return "", err
	}
	defer func() {
		err = tx.Rollback(ctx) // Откатываем транзакцию в случае ошибки.
		if err != nil {
			fmt.Printf("failed to rollback transaction: %v", err)
		}
	}()

	// Создаем экземпляр requestor для работы с продуктами.
	requestor := product.New(tx)
	requestor.WithTx(tx)

	// Получаем всех продуктов для проверки существования.
	products, err := requestor.GetAllProducts(ctx)
	if err != nil {
		fmt.Printf("failed to get all products: %v", err)
		return "", err
	}

	// Проверяем, существует ли продукт с таким именем.
	for _, product := range products {
		if product.Name == name {
			errorString := fmt.Sprintf("product with name %v already exists with id: %v", product.Name, product.ID)
			fmt.Println(errorString)
			err = errors.New(errorString)
			return "", err
		}
	}

	// Создаем новый продукт.
	productID, err := requestor.CreateProduct(ctx, product.CreateProductParams{
		Name:  name,
		Price: numeric,
	})
	if err != nil {
		fmt.Printf("failed to create product: %v", err)
		return "", err
	}

	// Коммитим транзакцию.
	err = tx.Commit(ctx)
	if err != nil {
		fmt.Printf("failed to commit transaction: %v", err)
		return "", err
	}
	return productID.String(), nil // Возвращаем ID нового продукта.
}
