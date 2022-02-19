import PySimpleGUI as sg
from window_layouts import *
from api_layer import *

import json #!todo temp during refactors, should not be needed after full implementation of api layer

# helpers
mlist_right_click_options = ['Copy', 'Paste', 'Select All', 'Cut']

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

def enable_enter_clicks(event, window):
	QT_ENTER_KEY1 = 'special 16777220'
	QT_ENTER_KEY2 = 'special 16777221'

	if event in ('\r', QT_ENTER_KEY1, QT_ENTER_KEY2):         # Check for ENTER key
		# go find element with Focus
		elem = window.find_element_with_focus()
		if elem is not None and elem.Type == sg.ELEM_TYPE_BUTTON:       # if it's a button element, click it
			elem.Click()

#windows
def change_pin_popup():
	layout = change_pin_window()
	window = sg.Window("Change Pin", layout, modal=True, return_keyboard_events=True)

	while True:
		event, values = window.read()
		enable_enter_clicks(event, window)

		if event == "Exit" or event == sg.WIN_CLOSED or event == 'none':
			break

		if event == '-SHOW_PIN-':
			if window['-SHOW_PIN-'].get():
				window['-PIN-'].update(password_char='')
			else:
				window['-PIN-'].update(password_char='*')
		
		if event == '-PIN-': #!todo This triggers on every char added so perfect place to check for restrictions on pin/pass
			print(window['-PIN-'].get())

		if event == '-APPLY_PIN-':
			print('Setting new pin to: ' + window['-PIN-'].get())
			break

		if event == '-CANCEL-': #!todo could go with exit but for now here in case I decide to do something extra with it.
			break

	window.close()

def delete_usr_popup():
	layout = delete_usr_window()
	window = sg.Window("Are you sure?", layout, modal=True, size=(245,40), return_keyboard_events=True)

	while True:
		event, values = window.read()
		enable_enter_clicks(event, window)

		if event in ("Exit", sg.WIN_CLOSED, 'none'):
			break

		if event == '-NO_DEL-':
			break

		if event == '-CONFIRM_DELETE-':
			break #!todo send delete request!

	window.close()

def main_window_event_loop(): #!todo functionize these events instead of bunch of stray code, for now hardcoded until using actual api calls
	data = load_dummy_dat()
	completeNameList = get_usr_data(data) #!todo, the ui lib adds a threading abstraction for slow/long operations (window.perform_long_operation) which could potentionally be used for async refresh, need to look better into async with the ui in general for this purpose.
	
	layout = main_window()
	window = sg.Window("KnT", layout, resizable=True, return_keyboard_events=True, finalize=True) #!todo, replace python logo with cute tiny storm logo?
	window['-USERLIST-'].update(completeNameList)

	while True:
		event, values = window.read() #!todo think about using elif for events, can there be more than one event per loop? lets assume no
		enable_enter_clicks(event, window)

		print('And now on this was the event, ' + str(event) + ' !')
		#print('values:')
		#print(values)

		if event  in ("Exit", sg.WIN_CLOSED, 'none'):
			break

		if event == '-USERLIST-' and len(values['-USERLIST-']) > 0:
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

		if event == '-CHANGE_PIN-':
			change_pin_popup()

		if event == '-DEL_USR-': #!todo confirmation pop up
			delete_usr_popup()

		if event == '-FILTER-':
			new_list = [i for i in completeNameList if (values['-FILTER-'].lower() in i.__str__().lower() or values['-FILTER-'].lower() in i.vunetId.lower())]
			window['-USERLIST-'].update(new_list)

		if event == '-ADD_USER-':
			add_user_popup()

		if event in mlist_right_click_options:
			do_clipboard_operation(event, window, window['-TRANSACTION_COMMENT-'])
		
		completeNameList = get_usr_data(data) #!todo async or something prolly

	window.close()

"""popup Window that handles adding a user"""
def add_user_popup():
	layout = add_usr_window()
	window = sg.Window("Add user", layout, modal=True, return_keyboard_events=True)
	mlist_right_click_options = ['Copy', 'Paste', 'Select All', 'Cut']
	mline:sg.Multiline = window['-TRANSACTION_COMMENT-']

	while True:
		event, values = window.read()
		enable_enter_clicks(event, window)

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
