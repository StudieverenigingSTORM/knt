from flask import Flask
from flask_cors import CORS
from config import Config
from flask_sqlalchemy import SQLAlchemy
from flask_migrate import Migrate
from flask_marshmallow import Marshmallow
from flask_restful import Api, Resource

app = Flask(__name__)
app.config.from_object(Config)
db = SQLAlchemy(app)
migrate = Migrate(app, db)
ma = Marshmallow(app)
api = Api(app)
cors = CORS(app, resources={r"/*": {"origins": "*"}})

from app import models

api.add_resource(models.ProductListResource, '/api/products')
api.add_resource(models.ProductResource, '/api/product/<int:id>')
api.add_resource(models.UsersListResource, '/api/users')
api.add_resource(models.UserResource, '/api/user/<int:id>')