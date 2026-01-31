FROM python:3.9-slim
WORKDIR /app

# Copy file requirements trước để tận dụng cache
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy source code
COPY app.py .

EXPOSE 5000
CMD ["python", "app.py"]
