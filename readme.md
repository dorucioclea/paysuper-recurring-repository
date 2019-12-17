PaySuper Repository
===

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-brightgreen.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Build Status](https://travis-ci.org/paysuper/paysuper-recurring-repository.svg?branch=master)](https://travis-ci.org/paysuper/paysuper-recurring-repository) 
[![codecov](https://codecov.io/gh/paysuper/paysuper-recurring-repository/branch/master/graph/badge.svg)](https://codecov.io/gh/paysuper/paysuper-recurring-repository)
[![Go Report Card](https://goreportcard.com/badge/github.com/paysuper/paysuper-recurring-repository)](https://goreportcard.com/report/github.com/paysuper/paysuper-recurring-repository)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=paysuper_paysuper-recurring-repository&metric=alert_status)](https://sonarcloud.io/dashboard?id=paysuper_paysuper-recurring-repository)

This service contain business logic for work with customer payment information. You can see service with name "paysuper-recurring-repository" on main architecture scheme by this [link](https://github.com/paysuper/paysuper-billing-server#architecture)

## Features

* Secure store customer payment information for simplify next payments
* Management method to find all saved customer payment methods
* Management method to find one saved customer payment method
* Management method to remove saved customer payment method 

## Developing

### Branches

We use the [GitFlow](https://nvie.com/posts/a-successful-git-branching-model) as a branching model for Git.

### Docker deployment

```bash
docker build -f Dockerfile -t paysuper_recurring_repository .
docker run -d -e "MONGO_DSN=mongodb://127.0.0.1:27017/recurring_repository" ... -e "METRICS_PORT=8081" paysuper_recurring_repository
```  

### Environment variables

|Name|Description|
|:---|:---|
| MONGO_DSN     | MongoDB DSN connection string                                                |
| MONGO_MODE    | Indicates the user's preference on reads from MongoDB. Default is "primary"  |
| METRICS_PORT  | Http server port for health and metrics request. Default is 8085             |

## Contributing
We feel that a welcoming community is important and we ask that you follow PaySuper's [Open Source Code of Conduct](https://github.com/paysuper/code-of-conduct/blob/master/README.md) in all interactions with the community.

PaySuper welcomes contributions from anyone and everyone. Please refer to each project's style and contribution guidelines for submitting patches and additions. In general, we follow the "fork-and-pull" Git workflow.

The master branch of this repository contains the latest stable release of this component.

 
