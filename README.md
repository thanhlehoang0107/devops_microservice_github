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
*(Hết bài 1 - Chuẩn bị code Python/Go)*
