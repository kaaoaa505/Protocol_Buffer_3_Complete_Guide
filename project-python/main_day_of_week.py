from build.PATH import dayofweek_pb2

message = dayofweek_pb2.TargetWeekObj()
message.id = 123
message.day_of_week = dayofweek_pb2.MONDAY

print(message)


