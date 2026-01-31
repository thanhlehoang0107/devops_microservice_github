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
*(Hết bài 2 - Tiếp theo sẽ là Docker)*
