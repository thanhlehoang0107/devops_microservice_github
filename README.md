# 📔 Nhật Ký DevOps Learning - GitHub Edition 🐙

Xin chào! Đây là repository mình thực hành DevOps trên nền tảng GitHub. Mình sẽ xây dựng hệ thống Microservices và dùng GitHub Actions để CI/CD.

---

## 📌 Bài 1: Thiết lập CI/CD với GitHub Actions

### 1. GitHub Actions khác gì GitLab CI?
Mình chuyển từ GitLab sang GitHub và nhận thấy sự khác biệt:
- Thay vì file `.gitlab-ci.yml`, GitHub dùng thư mục `.github/workflows/`.
- File cấu hình cũng là YAML nhưng cú pháp khác (dùng `jobs`, `steps`, `runs-on`).
- GitHub dùng "Runners" của Microsoft (ubuntu-latest) rất tiện.

### 2. Cách mình làm
Mình tạo file `.github/workflows/devops-pipeline.yml`. Mục tiêu là mỗi khi code được push lên nhánh `main`, pipeline sẽ tự chạy.

Cấu trúc pipeline của mình:
- **Trigger**: `on: push` (Khi có code mới).
- **Jobs**: Mình định nghĩa 3 job chạy tuần tự (dùng `needs` để job sau chờ job trước).

### 3. Cấu hình Demo
Tương tự bên GitLab, mình chạy thử một pipeline rỗng để test:

```yaml
name: DevOps Learning Pipeline
on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v3
      - name: Run Build
        run: echo "Building project..."

  test:
    needs: build
    runs-on: ubuntu-latest
    steps:
      - name: Run Test
        run: echo "Testing..."
```

### 4. Kết quả
- Tab **Actions** trên GitHub đã hiện tick xanh ✅.
- Mình đã hiểu cách dùng `needs` trong GitHub Actions để tạo dependency giữa các jobs (nếu không có `needs`, chúng sẽ chạy song song mặc định, khác với stage bên GitLab).


---

## 📌 Bài 2: Xây dựng Microservices (Python & Go)

### 1. Ý tưởng hệ thống
Hôm nay mình bắt tay vào coding. Hệ thống gồm 2 services đơn giản:
- **Go Backend**: Port 8080, xử lý logic nhanh.
- **Python Gateway**: Port 5000, nhận request từ user và gọi backend.

### 2. Code Implementation
Mình đã copy code `go-service` và `python-service` vào repo.
- **Python**: Dùng `requests` để gọi API. Cần xử lý `try-except` để tránh crash nếu Go service chưa bật.
- **Go**: Dùng `gorilla/mux` hoặc thư viện chuẩn. Ở đây mình dùng thư viện chuẩn cho đơn giản.

### 3. Cập nhật GitHub Actions
Mình sửa file `.github/workflows/devops-pipeline.yml`. Một điểm hay của GitHub Actions là `matrix strategy`. Mình có thể test code của mình trên nhiều version Python/Go cũng một lúc!

Nhưng để đơn giản cho bài này, mình tách làm 2 job:
- `test_go`: Setup Go environment -> Test.
- `test_python`: Setup Python environment -> Install reqs -> Lint.

### 4. Bài học rút ra
Việc cấu hình GitHub Actions (dùng `actions/setup-go`, `actions/setup-python`) cảm giác "thân thiện" hơn việc phải chọn Docker Image bên GitLab một chút, vì mình không cần quan tâm container bên dưới là gì.


---

## 📌 Bài 3: Docker Integration (Phần 1: Container hóa)

### 1. Ý tưởng
Đẩy code lên GitHub là một chuyện, nhưng người khác pull về có chạy được không? Có cần cài Python 3.9 hay 3.10? Cài Go version nào?
Giải pháp: **Docker**.

### 2. Implementation
Mình đã thêm `Dockerfile` vào từng thư mục service.
- **Go**: Sử dụng Multi-stage build để tối ưu dung lượng image.
- **Python**: Sử dụng slim image cho nhẹ.

Ngoài ra, file `docker-compose.yml` giúp mình định nghĩa toàn bộ stack. Chỉ cần `docker-compose up` là cả hệ thống backend + frontend (gateway) sẽ chạy lên.

### 3. Note về GitHub Actions
Hiện tại pipeline vẫn đang test code trần (không qua docker). Ở bài sau mình sẽ cập nhật pipeline để build và push docker image lên GitHub Packages (GHCR).


---

## 📌 Bài 3: Docker Integration (Phần 2: Networking & Ping)

### 1. Docker DNS
GitHub Repo này đã được cập nhật `docker-compose.yml` có cấu hình `networks`.
Các container giao tiếp qua tên service: `http://go-service:8080`.

### 2. Cập nhật Workflow
Mình đã thêm job `build_docker` vào GitHub Actions để đảm bảo `Dockerfile` không bị lỗi cú pháp trước khi merge code.


---

## 📌 Bài 3: Docker Integration (Phần 3: Full Feature CRUD)

### 1. Tính năng mới
Repo này giờ đã hoàn thiện tính năng:
- **Events API**: Cho phép Thêm/Sửa/Xóa events.
- **In-memory DB**: Go service lưu dữ liệu trong RAM.

### 2. Workflow hoàn chỉnh
Code Python -> Gọi sang Go.
Client chỉ cần giao tiếp với Python service (Gateway).

### 3. Next Steps
Bài tiếp theo, chúng ta sẽ đưa cái đống container này lên **Kubernetes** để quản lý cho "xịn" (tự động restart, scale, load balance).


---

## 📌 Bài 4: Kubernetes Basics

### 1. Manifest Files
Mình đã thêm thư mục `k8s/` chứa các file định nghĩa hạ tầng:
- `go-deployment.yaml`: Chạy 2 pod Backend.
- `python-deployment.yaml`: Chạy 1 pod Gateway, expose qua NodePort 30001.

### 2. Service Discovery trong K8s
Khác với Docker Compose dùng tên service trong config, K8s dùng đối tượng `Service`.
Khi mình tạo `kind: Service` tên là `go-service`, K8s sẽ tạo một DNS record nội bộ.
=> Code Python gọi `http://go-service:8080` vẫn hoạt động bình thường mà không cần sửa dòng code nào! Đây là vẻ đẹp của sự nhất quán giữa các môi trường.

### 3. CI/CD Update
Trên GitHub Actions, mình chưa setup cụm K8s thật để deploy. Nhưng mình có thể thêm bước "Linting" (kiểm tra cú pháp) cho các file YAML để tránh lỗi ngớ ngẩn.

---
*(Hết bài 4 - Tiếp theo: Monitoring)*
