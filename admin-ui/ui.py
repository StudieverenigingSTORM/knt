import PySimpleGUI as sg
import json
import pathlib
from stormer import Stormer
from window_layouts import *

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

def main_window(usrDataList):
	sg.theme('SystemDefaultForReal')
	userLayout = usr_window(usrDataList)
	productLayout = [[sg.Text('products')]]

	tabs = [[
		sg.TabGroup([
				[sg.Tab('Stormers', userLayout, key='-USERTAB-'), sg.Tab('Products', productLayout, key='-PRODTAB-')]
			], expand_x=True, expand_y=True)
	]]

	return sg.Window("Tabs", tabs, resizable=True)

def find_user(id, data): #temp function for dummy data, will have to be refactored based on how data is pulled on init and refreshed/loaded on user selection from remote db api
	for user in data:
		if user['_id'] == id:
			return user

	return None

"""Clip board is lost on closing program, tkinter limitation apparently, solvable by using other package if needed."""
def do_clipboard_operation(event, window, element):
	if event == 'Select All':
		element.Widget.selection_clear()
		element.Widget.tag_add('sel', '1.0', 'end')
	elif event == 'Copy':
		try:
			text = element.Widget.selection_get()
			window.TKroot.clipboard_clear()
			window.TKroot.clipboard_append(text)
		except:
			print('Nothing selected')
	elif event == 'Paste':
		element.Widget.insert(sg.tk.INSERT, window.TKroot.clipboard_get())
	elif event == 'Cut':
		try:
			text = element.Widget.selection_get()
			window.TKroot.clipboard_clear()
			window.TKroot.clipboard_append(text)
			element.update('')
		except:
			print('Nothing selected')

def add_user_popup():
	layout = add_usr_window()
	window = sg.Window("Add user", layout, modal=True)
	mlist_right_click_options = ['Copy', 'Paste', 'Select All', 'Cut']
	mline:sg.Multiline = window['-TRANSACTION_COMMENT-']

	while True:
		event, values = window.read()

		if event == "Exit" or event == sg.WIN_CLOSED or event == 'none':
			break

		if event == '-ADD_USER-': #!todo add error checking when less tired and more perceptive.
			dataOut = {}
			dataOut['firstname'] = window['-FIRSTNAME-'].get()
			dataOut['lastname'] = window['-LASTNAME-'].get()
			dataOut['vunetid'] = window['-VUNETID-'].get()
			balanceOp = window['-BALANCE_OPERAND-'].get()
			if balanceOp != '':  #naughty digits only! tooltip?!
				dataOut['transaction amount'] = balanceOp # wait on db design, check int/float (depening on db design) and what about sign, do we ever take money :(
				dataOut['transaction comment'] = window['-COMMENT-'].get() #this is all intentionally very crude, not gonna invest time until I have some insight in how we store money (float or int) and how comments/logs are gonna be handled!

			print('updating: ' + json.dumps(dataOut))
			break #!todo report back to user how we went perhaps?

		if event == '-CANCEL-': #!todo could go with exit but for now here in case I decide to do something extra with it.
			break

		if event in mlist_right_click_options:
			do_clipboard_operation(event, window, mline)


	window.close()

def event_loop(window, data): #!todo functionize these events instead of bunch of stray code, for now hardcoded until using actual api calls
	completeNameList = get_usr_data(data) #!todo, the ui lib adds a threading abstraction for slow/long operations (window.perform_long_operation) which could potentionally be used for async refresh, need to look better into async with the ui in general for this purpose.
	mlist_right_click_options = ['Copy', 'Paste', 'Select All', 'Cut']
	ulist_right_click_options = ['View', 'Delete']
	mline:sg.Multiline = window['-TRANSACTION_COMMENT-']

	while True:
		event, values = window.read() #!todo think about using elif for events, can there be more than one event per loop? lets assume no

		print('And now on this was the event, ' + str(event) + ' !')
		#print('values:')
		#print(values)

		if event == "Exit" or event == sg.WIN_CLOSED or event == 'none':
			break

		if event == '-USERLIST-':
			Stormer = values['-USERLIST-'][0]

			userInfo = find_user(Stormer.vunetId, data)

			if userInfo == None:
				continue

			window['-FIRSTNAME-'].update(userInfo['firstname'])
			window['-LASTNAME-'].update(userInfo['lastname'])
			window['-VUNETID-'].update(userInfo['_id'])
			window['-BALANCE-'].update(userInfo['balance'])
			window['-USR_INFO_PANEL-'].update(visible=True)

		if event == '-APPLY_CHANGES-':
			dataOut = {}
			dataOut['firstname'] = window['-FIRSTNAME-'].get()
			dataOut['lastname'] = window['-LASTNAME-'].get()
			dataOut['vunetid'] = window['-VUNETID-'].get()
			#data['balance'] = window['-BALANCE-'].get() #if keep, do int check, prolly should not be editable directly though

			print('updating: ' + json.dumps(dataOut))

		if event == '-TRANSACTION-':
			None #!todo start using api calls first

		if event == 'DEL_USR': #!todo
			None

		if event == '-FILTER-':
			new_list = [i for i in completeNameList if (values['-FILTER-'].lower() in i.__str__().lower() or values['-FILTER-'].lower() in i.vunetId.lower())]
			window['-USERLIST-'].update(new_list)

		if event == '-ADD_USER-':
			add_user_popup()

		if event in mlist_right_click_options:
			do_clipboard_operation(event, window, mline)
		
		completeNameList = get_usr_data(data) #!todo async or something prolly

	window.close()

def main():
	data = load_dummy_dat() #will have to figure out a good way of retrieving this, perhaps pull list of names and only get aditional info on selection of name
	usrData = get_usr_data(data)
	window = main_window(usrData)
	event_loop(window, data)
	print('bye')

if __name__ == "__main__": #python!!!!
	main()
