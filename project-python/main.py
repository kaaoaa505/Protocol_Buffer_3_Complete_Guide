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

print('---- Write to binary file')
print(message)
# wb: mean Write-Binary
with open('message.bin', 'wb') as f:
    bytesAsString = message.SerializeToString()
    print('---- Write to binary file bytesAsString is: \n\t', bytesAsString)
    f.write(bytesAsString)

print('\n\n---- Read from binary file')
# rb: mean Read-Binary
with open('message.bin', 'rb') as f:
    messageRead = f.read()
    print('---- Read from binary file messageRead is: \n\t', messageRead)
    messageReadString = sample_pb2.SampleObj().FromString(messageRead)
    print(messageReadString)

