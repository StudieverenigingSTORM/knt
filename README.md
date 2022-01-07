# The New and Improved K&T System

**Setup**
1. Make sure you have Python3, ``pip`` and ``npm`` installed.
2. Run ``pip install pipenv``. This is a python package manager and works (and looks) similar to what we are used to with ``npm``.
3. Navigate to ``/knt-backend`` and run ``pipenv --three``.
4. Run ``pipenv shell`` to start up the virtual environment. This starts up the virtual environment. It allows us to install python packages for this project that won't interfere with your global python packages and configs.
5. Run ``pipenv install`` to install all the packages in the Pipfile. 
6. Once all the packages are installed, run ``flask run`` to have the Flask app start up and run! You should see the url of the app, as well as some debugging input.
7. Once your app is all up and running, open another command line and navigate ``/knt-font``. 
8. Run ``npm install`` to install all the packages in the ``package.json``
9. Run ``npm run dev``. The front-end app should be up and running now.

**Notes**
In the front-end file ``nuxt.config.js`` and in the ``knt.py`` file, the two urls that are used are based on my machine. Usually, React/Vue.js/Nuxt.js uses port 3000 and flask uses port 5000, but if this isn't the case for you, just change the ports in the urls but _don't commit the changes_. 

There are still some things that need to be figured out on this fork if we decide to go through with Flask, those things being:
- How to use Swagger and SQLAlchemy to create the REST API

**Relevant documentation**
- The (best tutorial for Flask ever)[https://blog.miguelgrinberg.com/post/the-flask-mega-tutorial-part-i-hello-world]
- (Official Flask Documentation)[https://flask.palletsprojects.com/en/2.0.x/quickstart/]
- (Nuxt documentation)[https://nuxtjs.org/]

## Why replace it?

The old one is very insecure and pretty much impossible to maintain.

## Goals

* Be ez to use
* Be ez to maintain
* Be safe (we're dealing with money)

## MVP Goals

When buying snacks / soda / beer:

1. User goes in front of tablet thingy at the bar
2. User searches for their name and clicks on it
3. User enters password
4. User selects desired products and clicks buy.

When topping up account: User goes to stormwebshop.nl and buys credits.
Webshop calls API on backend to update the account.

## Extra ideas

* Ability to create account directly from the tablet
	* The account

## Tech stack

Backend: Python?

*Rationale*: Known by everybody, easy to piece together a web app backend with Flask.

Backend: Go?

*Rationale*: Very easy to learn, makes self contained executables, it's not Python, super nice HTTP libs

Frontend: HTML/JS/CSS

*Rationale*: Much easier to update the interface and adapt it to different needs.

## So what is this about?

The K&T system (I assume it stands for Koffie en Thee although tea is free) allows STORMers to purchase
edible and drinkable items such as sodas, coffee, beer, snacks and instant noodles. Because there is no cashier,
STORMers have an account with credits in the system that they subtract by themselves through the software
whenever they're buying something.

The system is composed of a user interface running on a computer at the bar, a server holding the backend, and
the STORM web shop which is used to buy credits.

Who benefits from it: all STORMers, BorrelCie

## How it should work, technically

It is going to be a web app that runs on the K&T thing in the StoKa. While this setup would technically allow
STORMers to use the app from the comfort of their own mobile phones / laptops, I (Tudor) consider that it
wouldn't be a good idea, because:

* Going to the actual K&T (tablet) physically signals to other people in the room that you're indeed paying for your stuff.
* I'm afraid of letting it on the public web (solution: only available in STORM LAN over wifi (which doesn't exist yet))?

This piece of software is going to be a Python (or Go?) executable that provides a REST API for doing business.
The client is going to be made of a bunch of HTML pages that call that API through JS.

There will also be an administrative API used also by the webshop to update the balance.

![Overview](./docs/overview.svg)
