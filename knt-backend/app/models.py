from app import db, ma, api
from flask_restful import Resource
from flask import request

class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    first_name = db.Column(db.String)
    last_name = db.Column(db.String)
    nickname = db.Column(db.String)
    password = db.Column(db.String)
    balance = db.Column(db.Numeric)

    def __repr__(self):
        return '<User {}>'.format(self.first_name)
    
""" PRODUCTS """
    
class Product(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.Text)
    price = db.Column(db.Numeric)
    
    def __repr__(self):
        return '<Product {}>'.format(self.product_name)

class ProductSchema(ma.Schema):
    class Meta:
        model = Product
        
product_schema = ProductSchema()
products_schema = ProductSchema(many=True)

class ProductListResource(Resource):
    def get(self):
        p = Product.query.all()
        return products_schema.dump(p)
    
    def post(self):
        p = Product(
            name=request.json['name'],
            price=request.json['price']
        )
        db.session.add(p)
        db.session.commit()
        return product_schema.dump(p)

class ProductResource(Resource):
    def patch(self, id):
        p = Product.query.get_or_404(id)
        
        if 'name' in request.json:
            p.name = request.json['name']
        if 'price' in request.json:
            p.price = request.json['price']
        
        db.session.commit()
        return product_schema.dump(p)
    
    def delete(self, id):
        p = Product.query.get_or_404(id)
        db.session.delete(p)
        db.session.commit()
        return '', 204
