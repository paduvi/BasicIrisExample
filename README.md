# Basic Iris Example

Tổng hợp những Example cơ bản khi sử dụng Iris Framework.

##### Điểm đáng chú ý:

- Tổ chức, sắp xếp thư mục, chia nhỏ file phù hợp cho mục đích làm REST API.
- Cơ chế Dispatcher giúp cải thiện khả năng scale, tăng hiệu năng khi gọi nhiều HTTP Request, RPC Request cùng 1 lúc. Tham khảo bài viết: http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
- Sử dụng [`fasthttp`](https://github.com/valyala/fasthttp) Client thay cho [`net/http`](https://golang.org/pkg/net/http/) Client
- Sử dụng [`Apache Thrift`](https://thrift.apache.org/tutorial/go) để triển khai RPC.
- Source code của Mock Remote Service: https://github.com/paduvi/MockRemoteService

### Index Controller:

- Hello World làm quen với Iris Framework.
- Route `"/ping"`: đóng vai trò API Gateway, thực hiện HTTP Request Ping tới Remote Service

### Todo Controller:

- API ứng dụng Todo List đơn giản
- Response trả về dưới dạng JSON. 
- Việc Xem/Thêm/Xóa được thực hiện thông qua Database (trong khuôn khổ thử nghiệm Example Iris cơ bản nên sử dụng tạm 1 Mock Repo thay vì dùng DB thật)
- Không sử dụng Dispatcher. Nên áp dụng Dispatcher giống như Message Controller để tối ưu hiệu năng, vì bản chất việc gửi command tới database cũng là 1 kết nối TCP

### Message Controller:

- API ứng dụng Message List đơn giản
- Response trả về dưới dạng JSON.
- Việc Xem/Thêm/Xóa được thực hiện thông qua lời gọi HTTP Request tới Remote Service. 
- Ở đây server đóng vai trò là 1 API Gateway, mọi job đều được điều phối thông qua Dispatcher.

### TODO-LIST:

- Bổ sung Example Iris với Apache Thrift.
- Bổ sung Example Redis và Long-polling Task.
