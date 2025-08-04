// Инициализация gRPC-сервера + роутинг
func main() {
	cfg := config.Load() // Читает .env
	log.Printf("Сервер запущен на порту %s", cfg.GRPCPort)
	grpcServer := grpc.NewServer()
	proto.RegisterMatcherServer(grpcServer, &MatcherService{})
	lis, _ := net.Listen("tcp", ":50051")
	grpcServer.Serve(lis)
}