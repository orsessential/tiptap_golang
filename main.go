package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Item struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderID     int       `json:"order_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Orders struct {
	ID           int       `json:"id"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Items        []Item    `json:"items"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "oliv12345678"
	dbname   = "order_assignments"
)

func connectDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createOrder(c *gin.Context, db *sql.DB) {
	var order Orders
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}()

	insertOrderSQL := `INSERT INTO orders (customer_name, ordered_at, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err = tx.QueryRow(insertOrderSQL, order.CustomerName, order.OrderedAt, time.Now(), time.Now()).Scan(&order.ID)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, item := range order.Items {
		insertItemSQL := `INSERT INTO items (name, description, quantity, order_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
		_, err := tx.Exec(insertItemSQL, item.Name, item.Description, item.Quantity, order.ID, time.Now(), time.Now())
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}

func getAllOrders(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var orders []Orders
	for rows.Next() {
		var order Orders
		err := rows.Scan(&order.ID, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func getOrderByID(c *gin.Context, db *sql.DB) {
	orderID := c.Param("id")

	var order Orders
	err := db.QueryRow("SELECT id, customer_name, ordered_at, created_at, updated_at FROM orders WHERE id = $1", orderID).
		Scan(&order.ID, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, err := db.Query("SELECT id, name, description, quantity FROM items WHERE order_id = $1", orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}
	order.Items = items

	c.JSON(http.StatusOK, order)
}

func updateOrder(c *gin.Context, db *sql.DB) {
	orderID := c.Param("id")

	var order Orders
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}()

	updateOrderSQL := `UPDATE orders SET customer_name=$1, ordered_at=$2, updated_at=$3 WHERE id=$4`
	_, err = tx.Exec(updateOrderSQL, order.CustomerName, order.OrderedAt, time.Now(), orderID)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	deleteItemsSQL := `DELETE FROM items WHERE order_id=$1`
	_, err = tx.Exec(deleteItemsSQL, orderID)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, item := range order.Items {
		insertItemSQL := `INSERT INTO items (name, description, quantity, order_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
		_, err := tx.Exec(insertItemSQL, item.Name, item.Description, item.Quantity, orderID, time.Now(), time.Now())
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

func deleteOrder(c *gin.Context, db *sql.DB) {
	orderID := c.Param("id")

	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}()

	rows, err := db.Query("SELECT * FROM orders WHERE order_id = $1", orderID)
	if rows == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	deleteItemsSQL := `DELETE FROM items WHERE order_id=$1`
	_, err = tx.Exec(deleteItemsSQL, orderID)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	deleteOrderSQL := `DELETE FROM orders WHERE id=$1`
	_, err = tx.Exec(deleteOrderSQL, orderID)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

func getAllOrdersHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		getAllOrders(c, db)
	}
}

func createOrderHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		createOrder(c, db)
	}
}

func getOrderByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		getOrderByID(c, db)
	}
}

func updateOrderHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		updateOrder(c, db)
	}
}

func deleteOrderHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		deleteOrder(c, db)
	}
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/orders", getAllOrdersHandler(db))
	router.POST("/orders", createOrderHandler(db))
	router.GET("/orders/:id", getOrderByIDHandler(db))
	router.PUT("/orders/:id", updateOrderHandler(db))
	router.DELETE("/orders/:id", deleteOrderHandler(db))
	router.Run(":8080")
}
