from django.db import models
from django.contrib.auth.models import AbstractUser


# Create your models here.

# package kntdb


# id INTEGER "id" INTEGER NOT NULL
# price NUMERIC "price" NUMERIC
# name TEXT "name" TEXT
# visibility INTEGER "visibility" INTEGER NOT NULL DEFAULT 1
# taxcategory INTEGER "taxcategory" INTEGER


class Product(models.Model):
    id = models.AutoField(primary_key=True, editable=False)
    price = models.IntegerField()
    name = models.TextField()
    visibility = models.BooleanField(default=True)
    taxcategory = models.ForeignKey("TaxCategory", on_delete=models.PROTECT)

    def __str__(self):
        return self.name


# id INTEGER "id" INTEGER NOT NULL
# first_name VARCHAR(120) "first_name" VARCHAR(120) NOT NULL
# last_name VARCHAR(120) "last_name" VARCHAR(120) NOT NULL
# vunetid VARCHAR(120) "vunetid" VARCHAR(120) NOT NULL UNIQUE
# password NUMERIC "password" NUMERIC
# balance NUMERIC "balance" NUMERIC NOT NULL DEFAULT 0
# visibility INTEGER "visibility" INTEGER NOT NULL DEFAULT 1


class User(AbstractUser):
    id = models.AutoField(primary_key=True, editable=False)
    first_name = models.CharField(max_length=120)
    last_name = models.CharField(max_length=120)
    vunetid = models.CharField(max_length=120, unique=True)
    knt_password = models.IntegerField()
    balance = models.IntegerField(default=0)
    visibility = models.BooleanField(default=True)

    def __str__(self):
        return self.vunetid


# id INTEGER "id" INTEGER NOT NULL UNIQUE
# user_id INTEGER "user_id" INTEGER NOT NULL
# starting_balance INTEGER "starting_balance" INTEGER NOT NULL
# delta_balance INTEGER "delta_balance" INTEGER NOT NULL
# final_balance INTEGER "final_balance" INTEGER NOT NULL
# receipt_id INTEGER "receipt_id" INTEGER NOT NULL
# ref TEXT "ref" TEXT


class Transaction(models.Model):
    id = models.AutoField(primary_key=True, editable=False)
    user = models.ForeignKey("User", on_delete=models.PROTECT)
    starting_balance = models.IntegerField()
    delta_balance = models.IntegerField()
    final_balance = models.IntegerField()
    receipt = models.ForeignKey("Receipt", on_delete=models.PROTECT)
    ref = models.TextField()

    def __str__(self):
        return self.ref


class TransactionItem(models.Model):
    transaction = models.ForeignKey("Transaction", on_delete=models.PROTECT)
    product = models.ForeignKey("Product", on_delete=models.PROTECT)
    quantity = models.IntegerField()
    price = models.IntegerField()
    tax = models.IntegerField()

    # primary key is composite of transaction and product id
    class Meta:
        unique_together = (("transaction_id", "product_id"),)

    def __str__(self):
        return self.transaction


# id INTEGER "id" INTEGER NOT NULL UNIQUE
# data TEXT "data" TEXT NOT NULL
# timestamp TEXT "timestamp" TEXT NOT NULL


class Receipt(models.Model):
    id = models.AutoField(primary_key=True, editable=False)
    data = models.TextField()
    timestamp = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return self.data


class TaxCategory(models.Model):
    id = models.AutoField(primary_key=True, editable=False)
    name = models.CharField(max_length=200)
    percentage = models.IntegerField()

    def __str__(self):
        return self.name
