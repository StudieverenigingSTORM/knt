from app import app
from flask import jsonify

@app.route('/hello')
def index():
    return jsonify({'greeting': 'Hello', 'greeting2': 'World!'})