from flask import Flask, jsonify
import requests
import os
import sys

app = Flask(__name__)

# Lấy URL của Go service từ biến môi trường
GO_SERVICE_URL = os.environ.get('GO_SERVICE_URL', 'http://localhost:8080')
print(f"Python Service configuration: GO_SERVICE_URL={GO_SERVICE_URL}", file=sys.stderr)

@app.route('/')
def home():
    return jsonify({"message": "Python Service is running 🐍"})

@app.route('/health')
def health():
    return "OK", 200

@app.route('/call-go')
def call_go():
    try:
        # Gọi sang Go service
        target_url = f"{GO_SERVICE_URL}/ping"
        print(f"Calling Go service at: {target_url}", file=sys.stderr)
        
        response = requests.get(target_url, timeout=3)
        return jsonify({
            "message": "Python gọi Go thành công!",
            "go_url": target_url,
            "go_response": response.json()
        })
    except Exception as e:
        return jsonify({
            "message": "Gọi Go thất bại!",
            "target": target_url,
            "error": str(e)
        }), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
