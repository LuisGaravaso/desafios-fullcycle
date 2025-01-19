package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	"cleanarch/configs"
	"cleanarch/internal/event"
	"cleanarch/internal/event/handler"
	"cleanarch/internal/infra/database"
	"cleanarch/internal/infra/graph"
	"cleanarch/internal/infra/grpc/pb"
	"cleanarch/internal/infra/grpc/service"
	"cleanarch/internal/infra/web"
	"cleanarch/internal/infra/web/webserver"
	"cleanarch/internal/usecase"
	"cleanarch/pkg/events"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	log.Printf("Loading Configs")
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	log.Printf("Connecting to Database")
	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()
	eventOrderHandler := handler.NewOrderHandler(rabbitMQChannel)

	eventDispatcher := events.NewEventDispatcher()
	for _, ev := range event.EventsToRegister {
		eventDispatcher.Register(ev, eventOrderHandler)
	}

	orderRepository := database.NewOrderRepository(db)

	log.Printf("Starting Servers")

	// WebServer
	webOrderHandler := web.NewWebOrderHandler(eventDispatcher, orderRepository)
	webServer := webserver.NewWebServer(configs.WebServerPort)
	webServer.AddHandler("/order", webOrderHandler.Create)
	webServer.AddHandler("/order/{id}", webOrderHandler.FindById)
	webServer.AddHandler("/orders", webOrderHandler.FindAll)
	log.Println("Starting web server on port", configs.WebServerPort)
	go webServer.Start()

	// gRPC Server
	grpcServer := grpc.NewServer()
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, eventDispatcher)
	createOrderService := service.NewCreateOrderService(*createOrderUseCase)
	pb.RegisterCreateOrderServiceServer(grpcServer, createOrderService)

	getOrderUseCase := usecase.NewGetOrderUseCase(orderRepository, eventDispatcher)
	getOrderService := service.NewGetOrderService(*getOrderUseCase)
	pb.RegisterGetOrderServiceServer(grpcServer, getOrderService)
	reflection.Register(grpcServer)

	log.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	// GraphQL Server
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		GetOrderUseCase:    *getOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	log.Println("Servers ready to use!")
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
