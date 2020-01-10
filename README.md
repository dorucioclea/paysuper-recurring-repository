# PaySuper Recurring Repository

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-brightgreen.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/paysuper/paysuper-recurring-repository/issues)
[![Build Status](https://travis-ci.com/paysuper/paysuper-recurring-repository.svg?branch=develop)](https://travis-ci.com/paysuper/paysuper-recurring-repository) 
[![codecov](https://codecov.io/gh/paysuper/paysuper-recurring-repository/branch/develop/graph/badge.svg)](https://codecov.io/gh/paysuper/paysuper-recurring-repository)
[![Go Report Card](https://goreportcard.com/badge/github.com/paysuper/paysuper-recurring-repository)](https://goreportcard.com/report/github.com/paysuper/paysuper-recurring-repository)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=paysuper_paysuper-recurring-repository&metric=alert_status)](https://sonarcloud.io/dashboard?id=paysuper_paysuper-recurring-repository)

PaySuper Recurring Repository contains the business logic to work with information about a customer's payment.

The main PaySuper architecture scheme with the `"paysuper-recurring-repository"` microservice in the `PCI DSS Data` section:

![The PaySuper architecture scheme](https://github.com/paysuper/paysuper-billing-server/blob/develop/docs/schema.png)

***

## Features

* Securely stores information about customer payments to simplify next payments.
* Finds all saved customer's payment methods.
* Finds one saved customer's payment method.
* Removes a saved customer's payment method.

## Table of Contents

- [Developing](#developing)
- [Contributing](#contributing-feature-requests-and-support)
- [License](#license)

## Developing

### Branches

We use the [GitFlow](https://nvie.com/posts/a-successful-git-branching-model) as a branching model for Git.

### Docker deployment

```bash
docker build -f Dockerfile -t paysuper_recurring_repository

docker run -d -e "MONGO_DSN=mongodb://127.0.0.1:27017/recurring_repository" -e "METRICS_PORT=8081" paysuper_recurring_repository
```

### Environment variables

|Name|Description|
|:---|:---|
| MONGO_DSN     | MongoDB DSN connection string                                                |
| MONGO_MODE    | Indicates the user's preference on reads from MongoDB. Default is "primary".  |
| METRICS_PORT  | Http server port for health and metrics request. Default is 8085.             |

## Contributing, Feature Requests and Support

If you like this project then you can put a ⭐ on it. It means a lot to us.

If you have an idea of how to improve PaySuper (or any of the product parts) or have general feedback, you're welcome to submit a [feature request](../../issues/new?assignees=&labels=&template=feature_request.md&title=).

Chances are, you like what we have already but you may require a custom integration, a special license or something else big and specific to your needs. We're generally open to such conversations.

If you have a question and can't find the answer yourself, you can [raise an issue](../../issues/new?assignees=&labels=&template=issue--support-request.md&title=I+have+a+question+about+<this+and+that>+%5BSupport%5D) and describe what exactly you're trying to do. We'll do our best to reply in a meaningful time.

We feel that a welcoming community is important and we ask that you follow PaySuper's [Open Source Code of Conduct](https://github.com/paysuper/code-of-conduct/blob/master/README.md) in all interactions with the community.

PaySuper welcomes contributions from anyone and everyone. Please refer to [our contribution guide to learn more](CONTRIBUTING.md).

## License

The project is available as open source under the terms of the [GPL v3 License](https://www.gnu.org/licenses/gpl-3.0).