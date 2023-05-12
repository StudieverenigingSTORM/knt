from django.contrib import admin

# Register your models here.

from .models import Product, User, Transaction, TaxCategory

admin.site.register(Product)
admin.site.register(User)
admin.site.register(Transaction)
admin.site.register(TaxCategory)
