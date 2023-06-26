from build.PATH import nested_pb2

EyeColorEnum = nested_pb2.NestedObj().EyeColor
address = nested_pb2.NestedObj().Address()
address.city = 'Dammam'
address.country = 'SA'
address.address_line_1 = '123 fake st.'

message = nested_pb2.NestedObj()
message.age = 34

message.eye_color = EyeColorEnum.BROWN

message.addresses.append(address)

print(message)
