# The New and Improved K&T System

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
