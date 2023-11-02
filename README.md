# Profit Calculator

Makes financial calculations due to given options.

## How to Install

#### Clone the repo

`git clone git@github.com:serhankk/profit-calculator.git`

#### Build project

`go build .`

#### READY TO RUN!

---

## How to Use

#### Options

First, need a `-option` parameter. Options are:

- change-percentage [buy, sell]
- profit-by-price [buy, amount, sell]
- price-after-change [buy, change]
- price-for-profit [buy, amount, profit]
- change-for-profit [buy, amount, profit]

Given in paranthesis words are required flags in order to make the calculation in preferred option.

#### Flags

As mentioned above, totally we have 5 flags and for each option, we use different flags.

I described detailed below section them.

---

## Examples

#### Calculate the change percentage due to entry and exit price

`$ ./profit-calculator -option change-percentage -buy 10 -sell 20`

`>>> Entry Price: 10 | Exit Price: 20 | Change: %100`

---

#### Calculate the profit price due to entry price, amount and exit price

`$ ./profit-calculator -option profit-by-price -buy 10 -amount 50 -sell 20`

`>>> Entry Price: 10 | Exit Price: 20 | Profit: 500`

---

#### Calculate the new price after change percentage

`$ ./profit-calculator -option price-after-change -buy 10 -change 25`

`>>> Entry Price: 10 | Change: %25 | New Price: 12.5`

---

#### Calculate target price due to the profit price wanted

`$ ./profit-calculator -option price-for-profit -buy 10 -amount 50 -profit 200`

`>>> Entry Price: 10 | Amount: 50 | Target Profit: 200 | Target Price: 14`

---

#### Calculate target change due to the profit price wanted

`$ ./profit-calculator -option price-for-profit -buy 10 -amount 50 -profit 200`

`>>> Entry Price: 10 | Amount: 50 | Target Profit: 200 | Target Change: %40`

---
