from flask import Flask, jsonify
import requests
import os

app = Flask(__name__)

# Lấy URL của Go service từ biến môi trường, mặc định là localhost:8080
GO_SERVICE_URL = os.environ.get('GO_SERVICE_URL', 'http://localhost:8080')

@app.route('/')
def home():
    return jsonify({"message": "Python Service is running 🐍"})

@app.route('/call-go')
def call_go():
    try:
        # Gọi sang Go service
        response = requests.get(f"{GO_SERVICE_URL}/ping", timeout=2)
        return jsonify({
            "message": "Python gọi Go thành công!",
            "go_response": response.json()
        })
    except Exception as e:
        return jsonify({
            "message": "Gọi Go thất bại!",
            "error": str(e)
        }), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
