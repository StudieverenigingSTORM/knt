from app import db, ma, api
from flask_restful import Resource
from flask import request

""" USERS """


class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    first_name = db.Column(db.Text)
    last_name = db.Column(db.Text)
    nickname = db.Column(db.Text)
    password = db.Column(db.String(64))
    vunet_id = db.Column(db.String(6))
    balance = db.Column(db.Text)

    def __repr__(self):
        return '<User {}>'.format(self.first_name)


class UserSchema(ma.Schema):
    class Meta:
        fields = ("first_name", "last_name", "nickname", "balance")
        model = User


user_schema = UserSchema()
users_schema = UserSchema(many=True)


class UsersListResource(Resource):
    def get(self):
        u = User.query.all()
        response = users_schema.dump(u)
        return response

    def post(self):
        u = User(first_name=request.json['first_name'], last_name=request.json['last_name'],
                 nickname=request.json['nickname'], password=request.json['password'], balance=request.json['balance'], vunet_id=request.json['vunet_id'])
        db.session.add(u)
        db.session.commit()
        return 'User Created', 204


class UserResource(Resource):
    def patch(self, id):
        u = User.get_or_404(id)
        
        if 'first_name' in request.json:
            u.first_name = request.json['first_name']
        if 'last_name' in request.json:
            u.last_name = request.json['last_name']
        if 'nickname' in request.json:
            u.nickname = request.json['nickname']
        if 'password' in request.json:
            u.password = request.json['password']
        if 'balance' in request.json:
            u.balance = request.json['balance']
        if 'vunet_id' in request.json:
            u.vunet_id = request.json['vunet_id']
        
        db.session.commit()
        return 'User patched', 200
    
    def delete(self, id):
        u = User.query.get_or_404(id)
        db.session.delete(u)
        db.session.commit()
        return 'User deleted', 200


""" PRODUCTS """


class Product(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.Text)
    price = db.Column(db.String)

    def __repr__(self):
        return '<Product {}>'.format(self.product_name)


class ProductSchema(ma.Schema):
    class Meta:
        fields = ("id", "name", "price")
        model = Product


product_schema = ProductSchema()
products_schema = ProductSchema(many=True)


class ProductListResource(Resource):
    def get(self):
        p = Product.query.all()
        response = products_schema.dump(p)
        return response

    def post(self):
        p = Product(
            name=request.json['name'],
            price=request.json['price']
        )
        db.session.add(p)
        db.session.commit()
        return 'Product created', 204


class ProductResource(Resource):
    def patch(self, id):
        p = Product.query.get_or_404(id)

        if 'name' in request.json:
            p.name = request.json['name']
        if 'price' in request.json:
            p.price = request.json['price']

        db.session.commit()
        return 'Patched', 200

    def delete(self, id):
        p = Product.query.get_or_404(id)
        db.session.delete(p)
        db.session.commit()
        return '', 204
