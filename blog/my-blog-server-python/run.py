import os

from src.app import create_app, create_sio

if __name__ == '__main__':
#   env_name = os.getenv('FLASK_ENV')
    env_name = 'development'
    app = create_app(env_name)
    sio = create_sio(app)
    sio.run(app, debug=True, port=9000)
 
    