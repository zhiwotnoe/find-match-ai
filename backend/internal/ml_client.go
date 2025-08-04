// gRPC-клиент для ML-сервиса
func GetMatches(userID string) ([]Match, error) {
	conn, _ := grpc.Dial("ml-service:50052")
	client := proto.NewMatcherClient(conn)
	resp, _ := client.GetMatches(context.Background(), &proto.UserRequest{UserId: userID})
	return resp.Matches, nil
}