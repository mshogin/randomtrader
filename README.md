# Random Trader 
### http://randomtrader.net

<a href="https://mshogin.com/RandomTrader">
  <img src="./assets/img/random-trader.png" style="width:800px;padding-left:20%">
</a>

[![Build Status](https://travis-ci.org/mshogin/RandomTrader.svg?branch=master)](https://travis-ci.org/mshogin/RandomTrader)
[![Release](https://img.shields.io/github/release/mshogin/RandomTrader.svg?style=flat)](https://github.com/mshogin/RandomTrader/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/mshogin/RandomTrader)](https://goreportcard.com/report/github.com/mshogin/RandomTrader)
[![Codecov](https://codecov.io/gh/mshogin/RandomTrader/branch/master/graph/badge.svg)](https://codecov.io/gh/mshogin/RandomTrader)
[![License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat)](https://github.com/mshogin/shogun/blob/master/LICENSE)

## Random trader is an algorithmic trading system based on the random events.

The minimum viable product contains the minimum acceptable functionality. 
The goal for MVP is to create a trading bot, 
integrate any neural network into the decision making service 
and run it with one exchange.

As cloud platform it proposed to use Heroku, 
since that service allows to setup initial infrastructure in short time.
Moreover, the actions in configuration of the CI/CD pipeline takes less time 
than with others services, especially, 
for the people without experience with cloud services, 
full-time at work and at home, and with a limited time in general. It is very convenient! ;)

The system consists of:
 - events generator service
 - feature management service
 - decision management service
 - exchange broker service
 - status service
 - metrics service
 
 The brief overview.
<img src="./assets/img/mvp.png" style="width:100%;padding-left:10%">

## Contribution

Please feel free to submit any pull requests.

## Contributor List


|User|
|--|
| [mshogin](https://github.com/mshogin) |

