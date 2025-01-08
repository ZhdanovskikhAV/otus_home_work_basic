package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw16_docker/internal/repository/transaction"
	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw16_docker/pkg/internal"
)

var DSN string

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	userData, err := transaction.GetUser(ctx, DSN, r.URL.Query().Get("name"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("Error Handled GET request for /v1/get_user with user: %+v", r.URL.Query().Get("name"))
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error,string,omitempty"`
		}{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(userData)
	w.WriteHeader(http.StatusOK)
	log.Printf("Handled GET request for /v1/get_user")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newUser internal.User
	ctx := context.Background()
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	userID, err := transaction.CreateUser(ctx, DSN, newUser.Name, newUser.Email, newUser.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error Handled POST request for /v1/create_user with user: %+v", newUser)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error,string,omitempty"`
		}{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		ID string `json:"id"`
	}{ID: userID})

	log.Printf("Handled POST request for /v1/create_user with user: %+v", newUser)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	orders, err := transaction.GetOrdersByUser(ctx, DSN, r.URL.Query().Get("name"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("Error Handled GET request for /v1/get_order with user: %+v", r.URL.Query().Get("name"))
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error,string,omitempty"`
		}{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
	log.Printf("Handled GET request for /v1/get_order")
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newOrder internal.Order
	ctx := context.Background()
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}
	totalAmount, ok := newOrder.TotalAmount.(float64)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	err = transaction.CreateOrder(ctx, DSN, newOrder.UserName, strconv.FormatFloat(totalAmount, 'f', 6, 64))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error Handled POST request for /v1/create_order with order: %+v", newOrder)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error,string,omitempty"`
		}{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	log.Printf("Handled POST request for /v1/create_order with order: %+v", newOrder)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()
	product, err := transaction.GetProductByName(ctx, DSN, r.URL.Query().Get("name"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("Error Handled GET request for /v1/get_product with product: %+v", r.URL.Query().Get("name"))
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error,string,omitempty"`
		}{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	log.Printf("Handled GET request for /v1/get_product")
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newProduct internal.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	ctx := context.Background()
	price, ok := newProduct.Price.(float64)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}
	_, err = transaction.CreateProduct(ctx, DSN, newProduct.Name, strconv.FormatFloat(price, 'f', 6, 64))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error Handled POST request for /v1/create_product with product: %+v", newProduct)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error,string,omitempty"`
		}{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	log.Printf("Handled POST request for /v1/create_product with product: %+v", newProduct)
}

func Server(ip *string, port *string, db *string) {
	// для получения через флаги
	// ip := flag.String("ip", "127.0.0.1", "IP address")
	// port := flag.String("port", "8080", "Port number")
	// flag.Parse()

	DSN = *db
	fmt.Printf("Starting server on %s:%s\n", *ip, *port)
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/v1/get_user", GetUser)
	serverMux.HandleFunc("/v1/create_user", CreateUser)
	serverMux.HandleFunc("/v1/get_order", GetOrder)
	serverMux.HandleFunc("/v1/create_order", CreateOrder)
	serverMux.HandleFunc("/v1/get_product", GetProduct)
	serverMux.HandleFunc("/v1/create_product", CreateProduct)

	server := &http.Server{
		Addr:         net.JoinHostPort(*ip, *port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      serverMux,
	}

	server.ListenAndServe()
}
