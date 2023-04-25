**MÔ TẢ**
- Ví dụ đơn giản về luồng chạy của app đặt vé, được viết theo kiểu microservice, chia các module thành cách service riêng biệt

- language: Golang
- database: Mongodb, elasticsearch
- framework: gRPC, Gin

**Cấu trúc thư mục**
- server: service phía backend, bao gồm customer và booking
- client: chứa các route, nhận dữ liệu đầu vào để truyền đến service phía dưới
- config: chứa địa chỉ config
- database_connection: khởi tạo kết nối đến database
- helpers: các hàm hỗ trợ
- protoc: chứa các định 
- pb: chứa các file interface service
- protoc: chứa các file định nghĩa model và service

+ gennerate code: dùng gitbash, chạy:
    - protoc --go_out=pb --go_opt=paths=import --go-grpc_out=pb --go-grpc_opt=paths=import protoc/tên_file.proto

+ sử dụng UI khi đã khởi chạy service:
    - grpcui -plaintext 127.0.0.1:port

Demo được chia thành 2 service: customer và booking
+ customer service bao gồm: tạo mới, sửa, xoá, thêm tag, xoá tag, tìm kiếm customer
+ booking service bao gồm: tạo ticket, tìm kiếm ticket

- phía server: bao gồm customer và booking và 2 server này chạy riêng biệt chịu trách nhiệm xử lí dữ liệu để trả về phía client, 2 service có thể giao tiếp với nhau, dù 1 service bị crash thì bên còn lại cũng không bị ảnh hưởng
- phía client: cũng bao gồm customer và booking tương ứng, chịu trách nhiệm nhận đầu vào dữ liệu từ phía frontend

**Cách chạy**
- "go run main.go" trong từng thưc mục của customer và booking service





