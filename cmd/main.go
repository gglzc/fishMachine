package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gglzc/fishMachine/grpc/server"
	"github.com/gglzc/fishMachine/pkg/mysql"
	pkg "github.com/gglzc/fishMachine/pkg/redis"
	pb "github.com/gglzc/fishMachine/proto"
	repositories "github.com/gglzc/fishMachine/repository/user"
	bulletServices "github.com/gglzc/fishMachine/service/bullet"
	userServices "github.com/gglzc/fishMachine/service/user"
	"github.com/gglzc/fishMachine/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 配置伺服器地址和端口
	address := ":50051"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Address Fail %v", err)
	}
	
	db, err := mysql.InitDB()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	
	redisClient, err := pkg.InitRedis() 
	if err != nil {
		log.Fatalf("Redis initialization failed: %v", err)
	}

	defer db.Close()

	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}
	log.Printf("Server listening on %s", address)
	//生成魚
	utils.LoadFishConfig()
	// 創建 gRPC 伺服器
	grpcServer := grpc.NewServer()
	// 初始化業務邏輯服務
	userRepository := &repositories.UserRepository{
		DB:    db,
		Redis: redisClient,
	}
	userService := &userServices.UserService{UserRepo: userRepository}
	bulletService := &bulletServices.GameService{
		UserService: userService,
	}
	
	// 註冊伺服器
	gameServer := &server.GameServer{
		BulltetService: bulletService,
		UserService:    userService,
	}
	pb.RegisterGameServiceServer(grpcServer, gameServer)

	// 啟用 gRPC 反射，方便開發和測試
	reflection.Register(grpcServer)

	// 使用 goroutine 启动 gRPC 服务器
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// 捕获系统信号以实现优雅关闭
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // 阻塞直到接收到终止信号
	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}
