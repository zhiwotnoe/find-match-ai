// Инициализация gRPC-сервера + роутинг
func main() {
	grpcServer := grpc.NewServer()
	proto.RegisterMatcherServer(grpcServer, &MatcherService{})
	lis, _ := net.Listen("tcp", ":50051")
	grpcServer.Serve(lis)
}