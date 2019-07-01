from flask import abort, request, Flask, jsonify
from flask_cors import CORS
from flask_socketio import SocketIO, emit
from .config import app_config

def create_app(env_name):
    app = Flask(__name__)
    app.config['SECRET_KEY'] = 'mysecret'
    app.config.from_object(app_config[env_name])
    CORS(app, supports_credentials=True)

    @app.route('/', methods=['GET'])
    def home():
        return jsonify({'status': "Worked"})

    @app.route('/test', methods=['GET'])
    def test():
        sio.emit('test', data="test")

    @app.route('/status/', methods=['GET'])
    def status():
        return jsonify({'status': "Worked"})

    return app

def create_sio(app):
    sio = SocketIO(app, async_mode='eventlet')
    
    @sio.on('emit_ping')
    def pong_response(message):
        sio.emit('pong', data="pong")

    return sio