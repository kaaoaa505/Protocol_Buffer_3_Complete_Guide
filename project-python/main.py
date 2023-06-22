from build.PATH import sample_pb2

message = sample_pb2.SampleObj()
message.id = 123
message.name = 'Khaled Allam'
message.is_valid = True

message.str_list.append('khaled')
message.str_list.append('allam')
message.str_list.append('ahmed')
message.str_list.append('omar')
message.str_list.append('abdalla')
message.str_list.append('allam')

print(message)

# wb: mean Write-Binary
with open('message.bin', 'wb') as f:
    bytesAsString = message.SerializeToString()
    f.write(bytesAsString)
