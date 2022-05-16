import json
import pathlib
from storageClasses import Stormer
from storageClasses import Product

#for now all mock calls on the dummy data, intention is to later abstract or atleast wrap all api calls through here

# !todo decide who ensures correctness, 
# does api caller pass whatever and possibly get an error back from here,
# is api caller solely repsonsible for making good calls
# or double work (both sides do their own checks), prolly this.

def load_dummy_data(file):
	filePath = pathlib.Path(pathlib.Path(__file__).parent, file).resolve()
	with open(filePath, 'r') as dummyboidatafile:
		dummyData = json.load(dummyboidatafile)

	print('dummy data contains ' + str(len(dummyData)) + ' entries.')
	return dummyData


def get_usr_data(inData): #works on temp dummy data, not sure what api calls I get yet!
	userList = []

	for user in inData: #!todo, if time for fun try a list comprehension instead!
		userList.append(Stormer(user['firstname'], user['lastname'], user['_id'], user['balance']))

	userList.sort(key=lambda x: str(x)) # performance note (in python, hilarious!) refactor to use operator instead of lambda might* be better 

	return userList

def get_prod_data(inData): #works on temp dummy data, not sure what api calls I get yet!
	productList = []

	for product in inData: #!todo, if time for fun try a list comprehension instead!
		productList.append(Product(product['name'], product['id'], product['price'], product['hidden']))

	productList.sort(key=lambda x: str(x)) # performance note (in python, hilarious!) refactor to use operator instead of lambda might* be better 

	return productList

def find_user(id, data): #temp function for dummy data, will have to be refactored based on how data is pulled on init and refreshed/loaded on user selection from remote db api
	for user in data:
		if user['_id'] == id:
			return user

	return None

def find_product(id, data): #temp function for dummy data, will have to be refactored based on how data is pulled on init and refreshed/loaded on user selection from remote db api
	for product in data:
		if product['id'] == id:
			return product

	return None

def add_user(firstName, lastName, vunetId, deposit, comment):
	dataOut = {}
	dataOut['firstname'] = firstName
	dataOut['lastname'] = lastName
	dataOut['vunetid'] = vunetId
	if deposit != '':  #digits only! tooltip?! Also always enforce a initial deposit even if 0?
		dataOut['transaction amount'] = deposit # wait on db design, check int/float (depening on db design) and what about sign, do we ever take money :(
		dataOut['transaction comment'] = comment #this is all intentionally very crude, not gonna invest time until I have some insight in how we store money (float or int) and how comments/logs are gonna be handled!

	print('updating: ' + json.dumps(dataOut))

def add_product(name, price, id):
	dataOut = {}
	dataOut['name'] = name
	dataOut['price'] = price
	dataOut['id'] = id

	print('updating: ' + json.dumps(dataOut))

def update_user(oldVunetId, firstName = None, lastName = None, newVunetId = None): #todo I hope the api will support dynamic/incomplete calls, if not refactor
	dataOut = {}	# old and new vunetId very depend on api design!
	print('updating user not rdy :(' + json.dumps(dataOut))

def del_user(vunetId):
	dataOut = {'vunetId' : vunetId}
	print('deleting id: ' + json.dumps(dataOut))

def del_product(id):
	dataOut = {'prodid' : id}
	print('deleting id: ' + json.dumps(dataOut))

def update_pin(vunetId, pin):
	dataOut = {'vunetId' : vunetId, 'pin' : pin}
	print('Setting new pin: ' + json.dumps(dataOut))

def transaction(vunetId, amount, comment):
	dataOut = {'vunetId' : vunetId, 'amount' : amount, 'comment' : comment}
	print('Performing transaction: ' + json.dumps(dataOut))
