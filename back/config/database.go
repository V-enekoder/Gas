package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/V-enekoder/GasManager/src/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPortStr := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	if dbHost == "" || dbPortStr == "" || dbUser == "" || dbName == "" {
		log.Fatal("Error: Faltan variables de entorno esenciales para la conexión a PostgreSQL (DB_HOST, DB_PORT, DB_USER, DB_NAME)")
	}

	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatalf("Error: DB_PORT inválido, debe ser un número: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
		dbSSLMode,
	)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	var gormErr error
	DB, gormErr = gorm.Open(postgres.Open(dsn), gormConfig)

	if gormErr != nil {
		log.Fatalf("Error al conectar con la base de datos PostgreSQL: %v", gormErr)
	}

	log.Println("Conectado exitosamente a la base de datos PostgreSQL!")

}

func SyncDB() {
	if DB == nil {
		log.Fatal("Error: La conexión a la base de datos no ha sido inicializada. Llama a ConnectDB() primero.")
		return
	}

	log.Println("Sincronizando modelos con la base de datos...")
	err := DB.AutoMigrate(
		// --- NIVEL 0: Tablas de Catálogo (no dependen de otras) ---
		// Estas tablas no tienen claves foráneas a otras tablas de nuestra aplicación.
		&schema.Municipality{},
		&schema.TypeCylinder{},
		&schema.OrderState{},
		&schema.PaymentState{},
		&schema.ReportType{},
		&schema.ReportState{},

		// --- NIVEL 1: Dependen del Nivel 0 ---
		&schema.User{}, // Depende de Municipality

		// --- NIVEL 2: Dependen del Nivel 1 ---
		&schema.Disabled{}, // Depende de User
		&schema.Council{},  // Depende de User
		&schema.Commerce{}, // Depende de User
		&schema.Order{},    // Depende de User y OrderState
		&schema.Payment{},  // Depende de PaymentState

		// --- NIVEL 3: Dependen del Nivel 2 ---
		// Esta es la parte crítica que corrige tu error.
		// Delivery DEBE estar antes que DeliveryDetail y Report.
		&schema.OrderDetail{}, // Depende de Order y TypeCylinder
		&schema.Delivery{},    // Depende de Order y Payment

		// --- NIVEL 4: Dependen del Nivel 3 ---
		&schema.DeliveryDetail{}, // Depende de Delivery
		&schema.Report{},         // Depende de Delivery
	)
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	log.Println("Modelos sincronizados exitosamente.")
}
