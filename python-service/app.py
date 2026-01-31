from flask import Flask, jsonify, request
import requests
import os
import sys

app = Flask(__name__)

GO_SERVICE_URL = os.environ.get('GO_SERVICE_URL', 'http://localhost:8080')

@app.route('/')
def home():
    return jsonify({"message": "Python Service Gateway 🐍"})

@app.route('/health')
def health():
    return "OK", 200

# Proxy endpoint để tạo Event
@app.route('/events', methods=['POST'])
def create_event():
    try:
        data = request.json
        # Forward request sang Go
        resp = requests.post(f"{GO_SERVICE_URL}/events", json=data, timeout=3)
        
        if resp.status_code == 201:
            return jsonify({
                "status": "success",
                "message": "Event created via Go",
                "data": resp.json()
            }), 201
        else:
            return jsonify({"status": "error", "go_error": resp.text}), resp.status_code
            
    except Exception as e:
        return jsonify({"status": "error", "error": str(e)}), 500

# Proxy endpoint để lấy Events
@app.route('/events', methods=['GET'])
def get_events():
    try:
        resp = requests.get(f"{GO_SERVICE_URL}/events", timeout=3)
        return jsonify({
            "status": "success",
            "source": "Go Service",
            "events": resp.json()
        })
    except Exception as e:
        return jsonify({"status": "error", "error": str(e)}), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
