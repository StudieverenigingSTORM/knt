import PySimpleGUI as sg
from window_layouts import *
import api_layer as api

"""Windows"""
def main_window_event_loop(): #!todo functionize these events instead of bunch of stray code, for now hardcoded until using actual api calls
	data = api.load_dummy_dat()
	completeNameList = api.get_usr_data(data) #!todo, the ui lib adds a threading abstraction for slow/long operations (window.perform_long_operation) which could potentionally be used for async refresh, need to look better into async with the ui in general for this purpose.
	
	layout = main_window_layout()
	window = sg.Window("KnT", layout, resizable=True, return_keyboard_events=True, finalize=True) #!todo, replace python logo with cute tiny storm logo?
	window['-USERLIST-'].update(completeNameList)
	startWidth = window.size[0]
	wideWidth = 0

	while True:
		event, values = window.read() #!todo think about using elif for events, can there be more than one event per loop? lets assume no
		enable_enter_clicks(event, window)
		print(event)

		if event  in ("Exit", sg.WIN_CLOSED, 'none'):
			break

		if event == '-USERLIST-' and len(values['-USERLIST-']) > 0: #!todo helper function, kinda bloats here
			Stormer = values['-USERLIST-'][0]

			userInfo = api.find_user(Stormer.vunetId, data)

			if userInfo == None:
				continue

			window['-FIRSTNAME-'].update(userInfo['firstname'])
			window['-LASTNAME-'].update(userInfo['lastname'])
			window['-VUNETID-'].update(userInfo['_id'])
			window['-BALANCE-'].update(userInfo['balance'])
			window['-USR_INFO_PANEL-'].update(visible=True)

		#todo, pretty disgusting, need to figure out how to automaticly make a window size fit visable content/layout if possible, rn works first time, but visable->hidden->visable doesnt!
			if wideWidth == 0:
				window.refresh() #sg is perfectly capable of picking a good fit and resizing on its own, perhaps follow .update(visable=True) to figure out why it works on initial call/what functions causes dynamic resize based on content, might happen in the window refresh as well!
				wideWidth = window.size[0]
			window.size = (wideWidth, window.size[1])

		if event == '-FILTER-':
			new_list = [i for i in completeNameList if (values['-FILTER-'].lower() in str(i).lower() or values['-FILTER-'].lower() in i.vunetId.lower())] #todo, hate the x in y or x in z, why wont x in (y or z) work properly!
			window['-USERLIST-'].update(new_list)

		if event == '-ADD_USER-':
			add_user_popup()

		if event == '-APPLY_CHANGES-':  #!todo, finish api call and refactor there and here
			dataOut = {}				#!todo Api design choice, possible to update vunetid or have to delete and add user for this?! #very important for this function, wait for api to take more form
			dataOut['firstname'] = window['-FIRSTNAME-'].get()
			dataOut['lastname'] = window['-LASTNAME-'].get()
			dataOut['vunetid'] = window['-VUNETID-'].get()

			print('updating: ' + str(dataOut))
			#api.update_user(...)

		if event == '-CHANGE_PIN-':
			#!todo super duper important, change this, should not read this from input since admin might just write in there by accident without clicking 'apply changes'
			change_pin_popup(window['-VUNETID-'].get()) 

		if event == '-DEL_USR-':
			fullName = window['-FIRSTNAME-'].get() + ' ' + window['-LASTNAME-'].get() #!todo same thing here, need to actually properly retrieve ID, just being lazy
			if delete_usr_popup(fullName, window['-VUNETID-'].get()):
				window['-USR_INFO_PANEL-'].update(visible=False)
				window['-FILTER-'].SetFocus()
				window.size = (startWidth, window.size[1]) #respect a users vertical resizes but force x size down their throat

		if event == '-TRANSACTION-':
			api.transaction(window['-VUNETID-'].get(), #!todo .
			window['-BALANCE_OPERAND-'].get(),
			window['-TRANSACTION_COMMENT-'].get()
			)


		if event in mlist_right_click_options:
			do_clipboard_operation(event, window, window['-TRANSACTION_COMMENT-'])
		
		completeNameList = api.get_usr_data(data) #!todo async or something prolly

	window.close()

def change_pin_popup(vunetId):
	layout = change_pin_window_layout()
	window = sg.Window("Change Pin", layout, modal=True, return_keyboard_events=True)

	while True:
		event, values = window.read()
		enable_enter_clicks(event, window)

		if event  in ("Exit", sg.WIN_CLOSED, 'none'):
			break

		if event == '-SHOW_PIN-':
			if window['-SHOW_PIN-'].get():
				window['-PIN-'].update(password_char='')
			else:
				window['-PIN-'].update(password_char='*')
		
		if event == '-PIN-': #!todo This triggers on every char added so perfect place to check for restrictions on pin/pass
			print(window['-PIN-'].get())

		if event == '-APPLY_PIN-':
			api.update_pin(vunetId, window['-PIN-'].get())
			break

		if event == '-CANCEL-': #!todo could go with exit but for now here in case I decide to do something extra with it.
			break

	window.close()

def delete_usr_popup(name, vunetId):
	layout = delete_usr_window_layout()
	window = sg.Window("Are you sure?", layout, modal=True, return_keyboard_events=True, finalize=True)
	msg = "You are about to remove " + name + " from the system!"
	window['-TEXT-'].update(msg) #window.read() will draw it in!

	while True:
		event, values = window.read()
		enable_enter_clicks(event, window)
		ret = False

		if event in ("Exit", sg.WIN_CLOSED, 'none', '-NO_DEL-'):
			break

		if event == '-CONFIRM_DELETE-':
			api.del_user(vunetId)
			ret = True
			break #!todo send delete request!

	window.close()
	return ret

def add_user_popup():
	layout = add_usr_window_layout()
	window = sg.Window("Add user", layout, modal=True, return_keyboard_events=True)

	while True:
		event, values = window.read()
		enable_enter_clicks(event, window)

		if event  in ("Exit", sg.WIN_CLOSED, 'none', '-CANCEL-'):
			break

		if event == '-ADD_USER-': #!todo add error checking when less tired and more perceptive.
			api.add_user(
				firstName = window['-FIRSTNAME-'].get(),
				lastName = window['-LASTNAME-'].get(),
				vunetId =  window['-VUNETID-'].get(),
				deposit = window['-BALANCE_OPERAND-'].get(),
				comment = window['-TRANSACTION_COMMENT-'].get()
			)
			break #!todo report back to user how we went perhaps?

		if event in mlist_right_click_options:
			do_clipboard_operation(event, window, window['-TRANSACTION_COMMENT-'])

	window.close()


"""Helpers"""
mlist_right_click_options = ['Copy', 'Paste', 'Select All', 'Cut']

"""Clip board is lost on closing program, tkinter limitation apparently, solvable by using other package for this if needed."""
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

	if event in ('\r', QT_ENTER_KEY1, QT_ENTER_KEY2):
		elem = window.find_element_with_focus()
		if elem is not None and elem.Type == sg.ELEM_TYPE_BUTTON:
			elem.Click()
