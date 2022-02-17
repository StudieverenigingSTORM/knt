from collections import UserList
from multiprocessing import dummy
import PySimpleGUI as sg
import json
import pathlib

filePath = pathlib.Path(pathlib.Path.cwd(), 'admin-ui/names.json')
with open(filePath, 'r') as dummyboidatafile:
    dummyData = json.load(dummyboidatafile)

print('my data len: ' + str(len(dummyData)))

userList = []
for user in dummyData:
    userList.append(user.get('name')) #there properly is a nicer way to do this instead of my noob loop, !todo

leftUsrCol = sg.Col([[sg.Listbox(values=userList, select_mode=sg.SELECT_MODE_EXTENDED, size=(20, 20), key='users')]])

#explicitly no middle part for name here for all non dutchies, if data base has names in first, middle, last merge middle and last.
personalInfo = sg.Col([
  [sg.Text('Personal info')],
  [sg.Text('Name:'), sg.Input(default_text = '', size=(25, 1), enable_events=True, key='-FIRSTNAME-')],
  [sg.Text('Last name:'), sg.Input(default_text = '', size=(25, 1), enable_events=True, key='-LASTNAME-')],
  [sg.Text('vunetID:'), sg.Input(default_text = '', size=(25, 1), enable_events=True, key='-VUNETID-')],
  ])

adminInfo = sg.Col([
  [sg.Text('Administrative')],
  [sg.Text('Balance:'), sg.Input(default_text = '', size=(25, 1), enable_events=True, key='-BALANCE-')],
  ])

rightUsrCol = sg.Column([[personalInfo], [adminInfo]])

userLayout = [[sg.Text('users', font='Any 16')],
              sg.vtop([leftUsrCol, rightUsrCol])
              ]

productLayout = [[sg.Text('products')]]

tabs = [[
    sg.TabGroup([[
        sg.Tab('users', userLayout), sg.Tab('products', productLayout)
        ]])
]]

#Define Window
window =sg.Window("Tabs", tabs, resizable=True, finalize=True)

#Read  values entered by user
while True:
    event, values = window.read(timeout = 5000, timeout_key = 1) # 5 sec update now for debug, crank up later !todo
    print('event')
    print(event)
    print('values')
    print(values)

    if event == sg.WIN_CLOSED:
        break

window.close()