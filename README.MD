## Snowflake - Unique ID generation library for Golang based on LockFree CAS algorithm

Snowflake is a powerful and easy-to-use library for generating unique IDs in your Golang applications. It uses a LockFree CAS algorithm to ensure thread-safety and prevent any possible collisions, making it an ideal choice for high-concurrency systems. The library provides a simple and intuitive interface to generate unique IDs with just a few lines of code.

### Features:

1. Thread-safe unique ID generation
2. LockFree CAS algorithm for high-concurrency systems
3. 4096 unique IDs per millisecond
4. Simple and intuitive interface
5. Easy to integrate with your existing Golang applications
6. Supports machine ID initialization using multiple methods

### Installation

To use Snowflake in your Golang project, you can use the go get command:

`go get github.com/dimail777/snowflake-go`

### Usage

Snowflake provides three initialization methods to generate unique IDs based on machine ID:

#### 1. InitByMachineId(machineId int64) (Snowflake, error)
This method initializes Snowflake with a specific machine ID. You can provide any 64-bit integer value as the
machine ID, as long as it's unique across your system.

#### 2. InitByRandom() (Snowflake, error)
This method initializes Snowflake with a randomly generated machine ID. The generated ID will be a 64-bit 
integer value that's unique across your system.

#### 3. InitByK8sStatefulSet() (Snowflake, error)
This method initializes Snowflake using the machine ID derived from the hostname of the Kubernetes StatefulSet. 
This method requires that the `HOSTNAME` environment variable is defined.

#### 4. Snowflake.GetNextId() (int64, error)
Once you've initialized Snowflake, you can generate unique IDs using the GetNextId() method.

### Contributing

We welcome contributions to Snowflake! To contribute, please fork this repository and submit a pull request. Before submitting your pull request, please ensure that your code follows the Go style guidelines and passes all tests.

### License

Snowflake is licensed under the Apache 2.0 license. See the LICENSE file for more information.

### Contributors

dimail777@gmail.com


## Links

* [SnowflakeID description](https://en.wikipedia.org/wiki/Snowflake_ID)