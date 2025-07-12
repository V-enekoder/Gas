package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/V-enekoder/GasManager/config"
	"github.com/V-enekoder/GasManager/src/schema"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	// 1. Cargar configuración y conectar a la base de datos
	log.Println("Iniciando programa de seeding...")
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()

	runSeeder("Parroquias/Municipios", SeedMunicipalities)
	runSeeder("Tipos de Cilindro", SeedTypeCylinders)
	runSeeder("Estados de Pedido", SeedOrderStates)
	runSeeder("Estados de Pago", SeedPaymentStates)
	runSeeder("Tipos de Reporte", SeedReportTypes)
	runSeeder("Estados de Reporte", SeedReportStates)
	runSeeder("Usuarios", SeedUsers)
	runSeeder("Ordenes", SeedOrders)
	runSeeder("Entregas", SeedDeliveries)

	log.Println("Programa de seeding finalizado exitosamente.")
}

func runSeeder(seederName string, seederFunc func(db *gorm.DB) error) {
	if err := seederFunc(config.DB); err != nil {
		log.Fatalf("Error durante el seeding de '%s': %v", seederName, err)
	}
}

// SeedMunicipalities puebla la tabla 'municipalities' con parroquias del Estado Bolívar.
func SeedMunicipalities(db *gorm.DB) error {
	municipalities := []schema.Municipality{
		// Municipio Caroní
		{Name: "Cachamay"},
		{Name: "Chirica"},
		{Name: "Dalla Costa"},
		{Name: "Once de Abril"},
		{Name: "Simón Bolívar"},
		{Name: "Unare"},
		{Name: "Universidad"},
		{Name: "Vista al Sol"},
		// Municipio Heres
		{Name: "Catedral"},
		{Name: "Agua Salada"},
		{Name: "La Sabanita"},
		{Name: "Marhuanta"},
		// Municipio Piar
		{Name: "Upata"},
	}

	for _, m := range municipalities {
		// FirstOrCreate evita crear duplicados si el seeder se corre múltiples veces.
		result := db.FirstOrCreate(&m, schema.Municipality{Name: m.Name})
		if result.Error != nil {
			return result.Error
		}
	}
	log.Println("Seeding de Parroquias/Municipios completado.")
	return nil
}

// SeedTypeCylinders puebla la tabla 'type_cylinder' con los tipos de bombonas de gas.
func SeedTypeCylinders(db *gorm.DB) error {
	cylinders := []schema.TypeCylinder{
		{Name: "Cilindro de 10kg", Description: "Bombona pequeña, ideal para familias o personas solas.", Price: 10.00, Disponible: true},
		{Name: "Cilindro de 18kg", Description: "Bombona mediana, de uso común en la mayoría de los hogares.", Price: 18.00, Disponible: true},
		{Name: "Cilindro de 43kg", Description: "Bombona grande, para uso comercial, restaurantes o familias numerosas.", Price: 43.00, Disponible: true},
	}

	for _, cylinder := range cylinders {
		result := db.FirstOrCreate(&cylinder, schema.TypeCylinder{Name: cylinder.Name})
		if result.Error != nil {
			return result.Error
		}
	}
	log.Println("Seeding de Tipos de Cilindro completado.")
	return nil
}

// SeedOrderStates puebla la tabla 'order_state' con los posibles estados de un pedido.
func SeedOrderStates(db *gorm.DB) error {
	states := []schema.OrderState{
		{Name: "Pendiente de Pago"},
		{Name: "En Proceso"},
		{Name: "En Ruta"},
		{Name: "Entregado"},
		{Name: "Completado"},
		{Name: "Verificado"},
		{Name: "Cancelado"},
	}

	for _, state := range states {
		result := db.FirstOrCreate(&state, schema.OrderState{Name: state.Name})
		if result.Error != nil {
			return result.Error
		}
	}
	log.Println("Seeding de Estados de Pedido completado.")
	return nil
}

// SeedPaymentStates puebla la tabla 'payment_state' con los posibles estados de un pago.
func SeedPaymentStates(db *gorm.DB) error {
	states := []schema.PaymentState{
		{Name: "Pendiente"},
		{Name: "Verificado"},
		{Name: "Rechazado"},
		{Name: "Reembolsado"},
	}

	for _, state := range states {
		result := db.FirstOrCreate(&state, schema.PaymentState{Name: state.Name})
		if result.Error != nil {
			return result.Error
		}
	}
	log.Println("Seeding de Estados de Pago completado.")
	return nil
}

// SeedReportTypes puebla la tabla 'report_type' con las categorías de incidencias.
func SeedReportTypes(db *gorm.DB) error {
	types := []schema.ReportType{
		{Name: "Cilindro dañado"},
		{Name: "Fuga de gas"},
		{Name: "Pedido incompleto"},
		{Name: "Problema con el pago"},
		{Name: "Otro"},
	}

	for _, rt := range types {
		result := db.FirstOrCreate(&rt, schema.ReportType{Name: rt.Name})
		if result.Error != nil {
			return result.Error
		}
	}
	log.Println("Seeding de Tipos de Reporte completado.")
	return nil
}

// SeedReportStates puebla la tabla 'report_state' con los estados de una incidencia.
func SeedReportStates(db *gorm.DB) error {
	states := []schema.ReportState{
		{Name: "Abierto"},
		{Name: "En Revisión"},
		{Name: "Resuelto"},
		{Name: "Cerrado"},
	}

	for _, state := range states {
		result := db.FirstOrCreate(&state, schema.ReportState{Name: state.Name})
		if result.Error != nil {
			return result.Error
		}
	}
	log.Println("Seeding de Estados de Reporte completado.")
	return nil
}

func SeedUsers(db *gorm.DB) error {
	// 1. Hashear la contraseña una sola vez para todos los usuarios de prueba
	password := "12345"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err // Si el hasheo falla, detenemos el seeder
	}
	hashedPasswordStr := string(hashedPassword)

	// --- Usuario 1: Perfil de Persona con Discapacidad ---
	userDisabled := schema.User{
		Name:           "Anaís Pérez (Discapacidad)",
		Email:          "anais.perez.disabled@example.com",
		Password:       hashedPasswordStr,
		MunicipalityID: 1, // Asumiendo que la Parroquia 'Cachamay' tiene ID 1
		Disabled: &schema.Disabled{
			Document:   "Certificado CONAPDIS 123456",
			Disability: "Movilidad reducida, requiere asistencia.",
		},
	}

	// --- Usuario 2: Perfil de Consejo Comunal ---
	userCouncil := schema.User{
		Name:           "Carlos Rodríguez (Jefe de Consejo)",
		Email:          "carlos.rodriguez.cc@example.com",
		Password:       hashedPasswordStr,
		MunicipalityID: 2, // Asumiendo que la Parroquia 'Chirica' tiene ID 2
		Council: &schema.Council{
			LeaderName:     "Carlos Rodríguez",
			LeaderDocument: "V-15123456",
		},
	}

	// --- Usuario 3: Perfil de Comercio ---
	userCommerce := schema.User{
		Name:           "María Fernández (Dueña de Restaurante)",
		Email:          "maria.fernandez.comercio@example.com",
		Password:       hashedPasswordStr,
		MunicipalityID: 3, // Asumiendo que la Parroquia 'Dalla Costa' tiene ID 3
		Commerce: &schema.Commerce{
			Rif:          "J-29123456-7",
			BossName:     "María Fernández",
			BossDocument: "V-12987654",
		},
	}

	usersToSeed := []schema.User{userDisabled, userCouncil, userCommerce}

	for _, user := range usersToSeed {
		// Creamos el usuario. GORM se encargará de crear el perfil asociado.
		// Buscamos por Email para evitar duplicados.
		result := db.FirstOrCreate(&user, schema.User{Email: user.Email})
		if result.Error != nil {
			return result.Error
		}
	}

	log.Println("Seeding de Usuarios y Perfiles completado.")
	return nil
}

func SeedOrders(db *gorm.DB) error {
	// --- 1. Obtener Datos Necesarios (Dependencias) ---

	// Obtener los usuarios por su email
	var commerceUser, councilUser, disabledUser schema.User
	db.First(&commerceUser, "email = ?", "maria.fernandez.comercio@example.com")
	db.First(&councilUser, "email = ?", "carlos.rodriguez.cc@example.com")
	db.First(&disabledUser, "email = ?", "anais.perez.disabled@example.com")

	if commerceUser.ID == 0 || councilUser.ID == 0 || disabledUser.ID == 0 {
		return errors.New("no se encontraron los usuarios de prueba. Ejecuta el seeder de usuarios primero")
	}

	// Obtener los tipos de cilindro
	var cylinder10kg, cylinder18kg, cylinder43kg schema.TypeCylinder
	db.First(&cylinder10kg, "name = ?", "Cilindro de 10kg")
	db.First(&cylinder18kg, "name = ?", "Cilindro de 18kg")
	db.First(&cylinder43kg, "name = ?", "Cilindro de 43kg")

	if cylinder10kg.ID == 0 || cylinder18kg.ID == 0 || cylinder43kg.ID == 0 {
		return errors.New("no se encontraron los tipos de cilindro. Ejecuta el seeder de cilindros primero")
	}

	// Obtener el estado inicial del pedido
	var pendingState schema.OrderState
	db.First(&pendingState, "name = ?", "Pendiente de Pago")
	if pendingState.ID == 0 {
		return errors.New("no se encontró el estado 'Pendiente de Pago'. Ejecuta el seeder de estados de pedido primero")
	}

	// --- 2. Crear las Órdenes de Prueba ---

	// Orden 1: El comercio pide varias bombonas grandes
	orderDetails1 := []schema.OrderDetail{
		{TypeCylinderID: cylinder43kg.ID, Quantity: 3, Price: cylinder43kg.Price},
		{TypeCylinderID: cylinder18kg.ID, Quantity: 2, Price: cylinder18kg.Price},
	}
	totalPrice1 := (float64(orderDetails1[0].Quantity) * orderDetails1[0].Price) + (float64(orderDetails1[1].Quantity) * orderDetails1[1].Price)

	order1 := schema.Order{
		UserID:       commerceUser.ID,
		TotalPrice:   totalPrice1,
		StateOrderID: pendingState.ID,
		OrderDetails: orderDetails1,
	}

	// Orden 2: El consejo comunal pide muchas bombonas pequeñas y medianas
	orderDetails2 := []schema.OrderDetail{
		{TypeCylinderID: cylinder10kg.ID, Quantity: 50, Price: cylinder10kg.Price},
		{TypeCylinderID: cylinder18kg.ID, Quantity: 30, Price: cylinder18kg.Price},
	}
	totalPrice2 := (float64(orderDetails2[0].Quantity) * orderDetails2[0].Price) + (float64(orderDetails2[1].Quantity) * orderDetails2[1].Price)

	order2 := schema.Order{
		UserID:       councilUser.ID,
		TotalPrice:   totalPrice2,
		StateOrderID: pendingState.ID,
		OrderDetails: orderDetails2,
	}

	// Orden 3: La persona con discapacidad pide una bombona pequeña
	orderDetails3 := []schema.OrderDetail{
		{TypeCylinderID: cylinder10kg.ID, Quantity: 1, Price: cylinder10kg.Price},
	}
	totalPrice3 := float64(orderDetails3[0].Quantity) * orderDetails3[0].Price

	order3 := schema.Order{
		UserID:       disabledUser.ID,
		TotalPrice:   totalPrice3,
		StateOrderID: pendingState.ID,
		OrderDetails: orderDetails3,
	}

	// --- 3. Guardar las órdenes en la base de datos ---
	ordersToSeed := []schema.Order{order1, order2, order3}
	for _, order := range ordersToSeed {
		result := db.Create(&order)
		if result.Error != nil {
			return result.Error
		}
	}

	log.Println("Seeding de Órdenes y Detalles de Orden completado.")
	return nil
}

func SeedDeliveries(db *gorm.DB) error {
	// --- 1. Obtener Dependencias ---

	// Buscar la orden del comercio, que es un buen caso de prueba.
	var commerceOrder schema.Order
	// Usamos Preload para cargar los OrderDetails junto con la orden.
	err := db.Preload("OrderDetails.TypeCylinder").Joins("User", db.Where(&schema.User{Email: "maria.fernandez.comercio@example.com"})).First(&commerceOrder).Error
	if err != nil {
		return errors.New("no se encontró la orden del comercio. Ejecuta los seeders anteriores")
	}

	// Verificar si esta orden ya tiene una entrega para no duplicar.
	var existingDelivery int64
	db.Model(&schema.Delivery{}).Where("order_id = ?", commerceOrder.ID).Count(&existingDelivery)
	if existingDelivery > 0 {
		log.Printf("La orden ID %d ya tiene una entrega. Se omite.", commerceOrder.ID)
		return nil
	}

	// Obtener estados necesarios
	var verifiedPaymentState schema.PaymentState
	var completedOrderState schema.OrderState
	db.First(&verifiedPaymentState, "name = ?", "Verificado")
	db.First(&completedOrderState, "name = ?", "Completado")
	if verifiedPaymentState.ID == 0 || completedOrderState.ID == 0 {
		return errors.New("no se encontraron los estados necesarios ('Verificado', 'Completado')")
	}

	// --- 2. Simular la Entrega Incompleta ---
	log.Printf("Procesando orden ID %d. Solicitado: %.2f", commerceOrder.ID, commerceOrder.TotalPrice)

	var deliveryDetails []schema.DeliveryDetail
	var deliveryTotalPrice float64 = 0.0

	for i, detail := range commerceOrder.OrderDetails {
		deliveredQuantity := detail.Quantity

		// SIMULACIÓN: Para el primer item del pedido, entregamos uno menos.
		if i == 0 && detail.Quantity > 1 {
			deliveredQuantity = detail.Quantity - 1
			log.Printf("  -> Item '%s': Solicitado %d, Entregado %d", detail.TypeCylinder.Name, detail.Quantity, deliveredQuantity)
		} else {
			log.Printf("  -> Item '%s': Solicitado %d, Entregado %d", detail.TypeCylinder.Name, detail.Quantity, deliveredQuantity)
		}

		deliveryDetails = append(deliveryDetails, schema.DeliveryDetail{
			TypeCylinderID: detail.TypeCylinderID,
			Quantity:       deliveredQuantity,
		})
		deliveryTotalPrice += float64(deliveredQuantity) * detail.Price
	}

	log.Printf("Valor real de la entrega: %.2f", deliveryTotalPrice)

	// --- 3. Crear Pago y Entrega en una Transacción ---
	err = db.Transaction(func(tx *gorm.DB) error {
		// a. Crear el Pago por el valor REAL de la entrega.
		payment := schema.Payment{
			UserID:   1,
			Quantity: deliveryTotalPrice,
			StateID:  verifiedPaymentState.ID,
		}
		if err := tx.Create(&payment).Error; err != nil {
			return err
		}

		// b. Crear la Entrega con sus detalles anidados.
		delivery := schema.Delivery{
			OrderID:         commerceOrder.ID,
			PaymentID:       payment.ID,
			TotalPrice:      deliveryTotalPrice,
			DeliveryDetails: deliveryDetails, // GORM creará estos detalles junto con la entrega.
		}
		if err := tx.Create(&delivery).Error; err != nil {
			return err
		}

		// c. Actualizar la orden a "Completado".
		if err := tx.Model(&commerceOrder).Update("StateOrderID", completedOrderState.ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	log.Println("Seeding de Entrega (incompleta) y Pago completado.")
	return nil
}

func SeedReports(db *gorm.DB) error {
	// --- 1. Obtener Dependencias ---

	// Buscar la entrega asociada al usuario de comercio.
	// El camino es: User -> Order -> Delivery.
	var targetDelivery schema.Delivery
	err := db.Preload("Order.OrderDetails.TypeCylinder"). // Preload para obtener los detalles originales
								Joins("JOIN orders ON orders.id = deliveries.order_id").
								Joins("JOIN users ON users.id = orders.user_id").
								Where("users.email = ?", "maria.fernandez.comercio@example.com").
								First(&targetDelivery).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("No se encontró la entrega del comercio. Seeding de Reportes omitido.")
			return nil
		}
		return errors.New("error al buscar la entrega del comercio")
	}

	// Verificar si ya existe un reporte para esta entrega.
	var existingReport int64
	db.Model(&schema.Report{}).Where("delivery_id = ?", targetDelivery.ID).Count(&existingReport)
	if existingReport > 0 {
		log.Printf("La entrega ID %d ya tiene un reporte. Se omite.", targetDelivery.ID)
		return nil
	}

	// Obtener el tipo de reporte y el estado inicial.
	var reportType schema.ReportType
	var reportState schema.ReportState
	db.First(&reportType, "name = ?", "Pedido incompleto")
	db.First(&reportState, "name = ?", "Abierto")

	if reportType.ID == 0 || reportState.ID == 0 {
		return errors.New("no se encontraron el tipo o el estado de reporte necesarios")
	}

	// --- 2. Construir y Crear el Reporte ---

	// Crear una descripción dinámica basada en la discrepancia.
	originalQuantity := targetDelivery.Order.OrderDetails[0].Quantity
	deliveredQuantity := int(targetDelivery.TotalPrice / targetDelivery.Order.OrderDetails[0].Price) // Asumiendo que solo un item fue entregado
	description := fmt.Sprintf(
		"Mi pedido original (Orden #%d) solicitaba %d cilindros de '%s', pero en la entrega solo recibí %d. El pago se realizó por el monto correcto, pero quiero dejar constancia de la discrepancia.",
		targetDelivery.OrderID,
		originalQuantity,
		targetDelivery.Order.OrderDetails[0].TypeCylinder.Name,
		deliveredQuantity,
	)

	report := schema.Report{
		DeliveryID:    targetDelivery.ID,
		Description:   description,
		Date:          time.Now(),
		TypeID:        reportType.ID,
		ReportStateID: reportState.ID,
	}

	if err := db.Create(&report).Error; err != nil {
		return err
	}

	log.Println("Seeding de Reporte por entrega incompleta completado.")
	return nil
}
