from app import db

class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    first_name = db.Column(db.String(120))
    last_name = db.Column(db.String(120))
    nickname = db.Column(db.String(120))
    password = db.Column(db.String(120))

    def __repr__(self):
        return '<User {}>'.format(self.first_name)
    
class Product(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    product_name = db.Column(db.Text)
    price = db.Column(db.Numeric)
    
    def __repr__(self):
        return '<Product {}>'.format(self.product_name)
