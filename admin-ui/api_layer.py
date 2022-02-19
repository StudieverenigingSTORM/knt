import json
import pathlib
from stormer import Stormer

#for now all mock calls on the dummy data, intention is to later abstract or atleast wrap all api calls through here

def load_dummy_dat():
	filePath = pathlib.Path(pathlib.Path(__file__).parent, 'names.json').resolve()
	with open(filePath, 'r') as dummyboidatafile:
		dummyData = json.load(dummyboidatafile)

	print('dummy data contains ' + str(len(dummyData)) + ' entries.')
	return dummyData

def get_usr_data(inData): #works on temp dummy data, not sure what api calls I get yet!
	userList = []

	for user in inData: #!todo, if time for fun try a list comprehension instead!
		userList.append(Stormer(user['firstname'], user['lastname'], user['_id'], user['balance']))

	userList.sort(key=lambda x: x.__str__()) # performance note (in python, hilarious!) refactor to use operator instead of lambda might* be better 

	return userList

def find_user(id, data): #temp function for dummy data, will have to be refactored based on how data is pulled on init and refreshed/loaded on user selection from remote db api
	for user in data:
		if user['_id'] == id:
			return user

	return None
